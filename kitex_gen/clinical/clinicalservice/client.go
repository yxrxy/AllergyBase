// Code generated by Kitex v0.13.1. DO NOT EDIT.

package clinicalservice

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
	clinical "github.com/yxrxy/AllergyBase/kitex_gen/clinical"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	CreatePatient(ctx context.Context, req *clinical.CreatePatientRequest, callOptions ...callopt.Option) (r *clinical.CreatePatientResponse, err error)
	GetPatient(ctx context.Context, req *clinical.GetPatientRequest, callOptions ...callopt.Option) (r *clinical.GetPatientResponse, err error)
	ListPatients(ctx context.Context, req *clinical.ListPatientsRequest, callOptions ...callopt.Option) (r *clinical.ListPatientsResponse, err error)
	CreateVisit(ctx context.Context, req *clinical.CreateVisitRequest, callOptions ...callopt.Option) (r *clinical.CreateVisitResponse, err error)
	GetVisit(ctx context.Context, req *clinical.GetVisitRequest, callOptions ...callopt.Option) (r *clinical.GetVisitResponse, err error)
	GetPatientVisits(ctx context.Context, req *clinical.GetPatientVisitsRequest, callOptions ...callopt.Option) (r *clinical.GetPatientVisitsResponse, err error)
	ListAllVisits(ctx context.Context, req *clinical.ListAllVisitsRequest, callOptions ...callopt.Option) (r *clinical.ListAllVisitsResponse, err error)
	AddDiagnosis(ctx context.Context, req *clinical.AddDiagnosisRequest, callOptions ...callopt.Option) (r *clinical.AddDiagnosisResponse, err error)
	GetVisitDiagnoses(ctx context.Context, req *clinical.GetVisitDiagnosesRequest, callOptions ...callopt.Option) (r *clinical.GetVisitDiagnosesResponse, err error)
	ListAllDiagnoses(ctx context.Context, req *clinical.ListAllDiagnosesRequest, callOptions ...callopt.Option) (r *clinical.ListAllDiagnosesResponse, err error)
	AddExamination(ctx context.Context, req *clinical.AddExaminationRequest, callOptions ...callopt.Option) (r *clinical.AddExaminationResponse, err error)
	GetVisitExaminations(ctx context.Context, req *clinical.GetVisitExaminationsRequest, callOptions ...callopt.Option) (r *clinical.GetVisitExaminationsResponse, err error)
	ListAllExaminations(ctx context.Context, req *clinical.ListAllExaminationsRequest, callOptions ...callopt.Option) (r *clinical.ListAllExaminationsResponse, err error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfoForClient(), options...)
	if err != nil {
		return nil, err
	}
	return &kClinicalServiceClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kClinicalServiceClient struct {
	*kClient
}

func (p *kClinicalServiceClient) CreatePatient(ctx context.Context, req *clinical.CreatePatientRequest, callOptions ...callopt.Option) (r *clinical.CreatePatientResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CreatePatient(ctx, req)
}

func (p *kClinicalServiceClient) GetPatient(ctx context.Context, req *clinical.GetPatientRequest, callOptions ...callopt.Option) (r *clinical.GetPatientResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetPatient(ctx, req)
}

func (p *kClinicalServiceClient) ListPatients(ctx context.Context, req *clinical.ListPatientsRequest, callOptions ...callopt.Option) (r *clinical.ListPatientsResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ListPatients(ctx, req)
}

func (p *kClinicalServiceClient) CreateVisit(ctx context.Context, req *clinical.CreateVisitRequest, callOptions ...callopt.Option) (r *clinical.CreateVisitResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CreateVisit(ctx, req)
}

func (p *kClinicalServiceClient) GetVisit(ctx context.Context, req *clinical.GetVisitRequest, callOptions ...callopt.Option) (r *clinical.GetVisitResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetVisit(ctx, req)
}

func (p *kClinicalServiceClient) GetPatientVisits(ctx context.Context, req *clinical.GetPatientVisitsRequest, callOptions ...callopt.Option) (r *clinical.GetPatientVisitsResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetPatientVisits(ctx, req)
}

func (p *kClinicalServiceClient) ListAllVisits(ctx context.Context, req *clinical.ListAllVisitsRequest, callOptions ...callopt.Option) (r *clinical.ListAllVisitsResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ListAllVisits(ctx, req)
}

func (p *kClinicalServiceClient) AddDiagnosis(ctx context.Context, req *clinical.AddDiagnosisRequest, callOptions ...callopt.Option) (r *clinical.AddDiagnosisResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.AddDiagnosis(ctx, req)
}

func (p *kClinicalServiceClient) GetVisitDiagnoses(ctx context.Context, req *clinical.GetVisitDiagnosesRequest, callOptions ...callopt.Option) (r *clinical.GetVisitDiagnosesResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetVisitDiagnoses(ctx, req)
}

func (p *kClinicalServiceClient) ListAllDiagnoses(ctx context.Context, req *clinical.ListAllDiagnosesRequest, callOptions ...callopt.Option) (r *clinical.ListAllDiagnosesResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ListAllDiagnoses(ctx, req)
}

func (p *kClinicalServiceClient) AddExamination(ctx context.Context, req *clinical.AddExaminationRequest, callOptions ...callopt.Option) (r *clinical.AddExaminationResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.AddExamination(ctx, req)
}

func (p *kClinicalServiceClient) GetVisitExaminations(ctx context.Context, req *clinical.GetVisitExaminationsRequest, callOptions ...callopt.Option) (r *clinical.GetVisitExaminationsResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetVisitExaminations(ctx, req)
}

func (p *kClinicalServiceClient) ListAllExaminations(ctx context.Context, req *clinical.ListAllExaminationsRequest, callOptions ...callopt.Option) (r *clinical.ListAllExaminationsResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ListAllExaminations(ctx, req)
}
