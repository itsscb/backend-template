package gapi

import (
	"context"
	"database/sql"
	"errors"
	"log/slog"

	db "github.com/itsscb/backend-template/bff/db/sqlc"
	"github.com/itsscb/backend-template/bff/pb"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) UpdateAccountPrivacy(ctx context.Context, req *pb.UpdateAccountPrivacyRequest) (*pb.UpdateAccountPrivacyResponse, error) {
	authPayload, err := server.authorizeUser(ctx)
	if err != nil {
		return nil, unauthenticatedError(err)
	}

	violations := validateUpdateAccountPrivacyRequest(req)
	if violations != nil {
		return nil, invalidArgumentError(violations)
	}

	account, err := server.store.GetAccount(ctx, req.GetId())
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, status.Errorf(codes.NotFound, "account not found")
		}
		slog.Error("update_account_privacy (get_account)", slog.Int64("invoked_by", int64(authPayload.AccountID)), slog.Int64("account_id", int64(req.GetId())), slog.String("error", err.Error()))
		return nil, status.Errorf(codes.Internal, "failed to get account")
	}

	if authPayload.AccountID != account.ID {
		if !server.isAdmin(ctx, authPayload) {
			return nil, status.Error(codes.NotFound, "account not found")
		}
	}
	privacyAccepted := req.GetPrivacyAccepted()

	arg := db.UpdateAccountPrivacyTxParams{
		Changer:         account.Email,
		ID:              req.GetId(),
		PrivacyAccepted: &privacyAccepted,
	}

	account, err = server.store.UpdateAccountPrivacyTx(ctx, arg)
	if err != nil {
		slog.Error("update_account_privacy (db)", slog.Int64("invoked_by", int64(authPayload.AccountID)), slog.Int64("account_id", int64(req.GetId())), slog.String("error", err.Error()))
		return nil, status.Error(codes.Internal, "failed to update account privacy")
	}

	rsp := &pb.UpdateAccountPrivacyResponse{
		Account: convertAccount(account),
	}

	return rsp, nil
}

func validateUpdateAccountPrivacyRequest(req *pb.UpdateAccountPrivacyRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if req.GetId() < 1 {
		violations = append(violations, fieldViolation("id", errors.New("must be greater than 0")))
	}
	return violations
}
