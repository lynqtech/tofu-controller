// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: runner/runner.proto

package runner

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// RunnerClient is the client API for Runner service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RunnerClient interface {
	LookPath(ctx context.Context, in *LookPathRequest, opts ...grpc.CallOption) (*LookPathReply, error)
	NewTerraform(ctx context.Context, in *NewTerraformRequest, opts ...grpc.CallOption) (*NewTerraformReply, error)
	SetEnv(ctx context.Context, in *SetEnvRequest, opts ...grpc.CallOption) (*SetEnvReply, error)
	CreateFileMappings(ctx context.Context, in *CreateFileMappingsRequest, opts ...grpc.CallOption) (*CreateFileMappingsReply, error)
	UploadAndExtract(ctx context.Context, in *UploadAndExtractRequest, opts ...grpc.CallOption) (*UploadAndExtractReply, error)
	CleanupDir(ctx context.Context, in *CleanupDirRequest, opts ...grpc.CallOption) (*CleanupDirReply, error)
	WriteBackendConfig(ctx context.Context, in *WriteBackendConfigRequest, opts ...grpc.CallOption) (*WriteBackendConfigReply, error)
	ProcessCliConfig(ctx context.Context, in *ProcessCliConfigRequest, opts ...grpc.CallOption) (*ProcessCliConfigReply, error)
	GenerateVarsForTF(ctx context.Context, in *GenerateVarsForTFRequest, opts ...grpc.CallOption) (*GenerateVarsForTFReply, error)
	Plan(ctx context.Context, in *PlanRequest, opts ...grpc.CallOption) (*PlanReply, error)
	ShowPlanFileRaw(ctx context.Context, in *ShowPlanFileRawRequest, opts ...grpc.CallOption) (*ShowPlanFileRawReply, error)
	ShowPlanFile(ctx context.Context, in *ShowPlanFileRequest, opts ...grpc.CallOption) (*ShowPlanFileReply, error)
	SaveTFPlan(ctx context.Context, in *SaveTFPlanRequest, opts ...grpc.CallOption) (*SaveTFPlanReply, error)
	LoadTFPlan(ctx context.Context, in *LoadTFPlanRequest, opts ...grpc.CallOption) (*LoadTFPlanReply, error)
	Apply(ctx context.Context, in *ApplyRequest, opts ...grpc.CallOption) (*ApplyReply, error)
	GetInventory(ctx context.Context, in *GetInventoryRequest, opts ...grpc.CallOption) (*GetInventoryReply, error)
	Destroy(ctx context.Context, in *DestroyRequest, opts ...grpc.CallOption) (*DestroyReply, error)
	Output(ctx context.Context, in *OutputRequest, opts ...grpc.CallOption) (*OutputReply, error)
	WriteOutputs(ctx context.Context, in *WriteOutputsRequest, opts ...grpc.CallOption) (*WriteOutputsReply, error)
	GetOutputs(ctx context.Context, in *GetOutputsRequest, opts ...grpc.CallOption) (*GetOutputsReply, error)
	Init(ctx context.Context, in *InitRequest, opts ...grpc.CallOption) (*InitReply, error)
	SelectWorkspace(ctx context.Context, in *WorkspaceRequest, opts ...grpc.CallOption) (*WorkspaceReply, error)
	Upload(ctx context.Context, in *UploadRequest, opts ...grpc.CallOption) (*UploadReply, error)
	FinalizeSecrets(ctx context.Context, in *FinalizeSecretsRequest, opts ...grpc.CallOption) (*FinalizeSecretsReply, error)
	ForceUnlock(ctx context.Context, in *ForceUnlockRequest, opts ...grpc.CallOption) (*ForceUnlockReply, error)
}

type runnerClient struct {
	cc grpc.ClientConnInterface
}

func NewRunnerClient(cc grpc.ClientConnInterface) RunnerClient {
	return &runnerClient{cc}
}

func (c *runnerClient) LookPath(ctx context.Context, in *LookPathRequest, opts ...grpc.CallOption) (*LookPathReply, error) {
	out := new(LookPathReply)
	err := c.cc.Invoke(ctx, "/runner.Runner/LookPath", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *runnerClient) NewTerraform(ctx context.Context, in *NewTerraformRequest, opts ...grpc.CallOption) (*NewTerraformReply, error) {
	out := new(NewTerraformReply)
	err := c.cc.Invoke(ctx, "/runner.Runner/NewTerraform", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *runnerClient) SetEnv(ctx context.Context, in *SetEnvRequest, opts ...grpc.CallOption) (*SetEnvReply, error) {
	out := new(SetEnvReply)
	err := c.cc.Invoke(ctx, "/runner.Runner/SetEnv", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *runnerClient) CreateFileMappings(ctx context.Context, in *CreateFileMappingsRequest, opts ...grpc.CallOption) (*CreateFileMappingsReply, error) {
	out := new(CreateFileMappingsReply)
	err := c.cc.Invoke(ctx, "/runner.Runner/CreateFileMappings", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *runnerClient) UploadAndExtract(ctx context.Context, in *UploadAndExtractRequest, opts ...grpc.CallOption) (*UploadAndExtractReply, error) {
	out := new(UploadAndExtractReply)
	err := c.cc.Invoke(ctx, "/runner.Runner/UploadAndExtract", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *runnerClient) CleanupDir(ctx context.Context, in *CleanupDirRequest, opts ...grpc.CallOption) (*CleanupDirReply, error) {
	out := new(CleanupDirReply)
	err := c.cc.Invoke(ctx, "/runner.Runner/CleanupDir", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *runnerClient) WriteBackendConfig(ctx context.Context, in *WriteBackendConfigRequest, opts ...grpc.CallOption) (*WriteBackendConfigReply, error) {
	out := new(WriteBackendConfigReply)
	err := c.cc.Invoke(ctx, "/runner.Runner/WriteBackendConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *runnerClient) ProcessCliConfig(ctx context.Context, in *ProcessCliConfigRequest, opts ...grpc.CallOption) (*ProcessCliConfigReply, error) {
	out := new(ProcessCliConfigReply)
	err := c.cc.Invoke(ctx, "/runner.Runner/ProcessCliConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *runnerClient) GenerateVarsForTF(ctx context.Context, in *GenerateVarsForTFRequest, opts ...grpc.CallOption) (*GenerateVarsForTFReply, error) {
	out := new(GenerateVarsForTFReply)
	err := c.cc.Invoke(ctx, "/runner.Runner/GenerateVarsForTF", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *runnerClient) Plan(ctx context.Context, in *PlanRequest, opts ...grpc.CallOption) (*PlanReply, error) {
	out := new(PlanReply)
	err := c.cc.Invoke(ctx, "/runner.Runner/Plan", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *runnerClient) ShowPlanFileRaw(ctx context.Context, in *ShowPlanFileRawRequest, opts ...grpc.CallOption) (*ShowPlanFileRawReply, error) {
	out := new(ShowPlanFileRawReply)
	err := c.cc.Invoke(ctx, "/runner.Runner/ShowPlanFileRaw", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *runnerClient) ShowPlanFile(ctx context.Context, in *ShowPlanFileRequest, opts ...grpc.CallOption) (*ShowPlanFileReply, error) {
	out := new(ShowPlanFileReply)
	err := c.cc.Invoke(ctx, "/runner.Runner/ShowPlanFile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *runnerClient) SaveTFPlan(ctx context.Context, in *SaveTFPlanRequest, opts ...grpc.CallOption) (*SaveTFPlanReply, error) {
	out := new(SaveTFPlanReply)
	err := c.cc.Invoke(ctx, "/runner.Runner/SaveTFPlan", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *runnerClient) LoadTFPlan(ctx context.Context, in *LoadTFPlanRequest, opts ...grpc.CallOption) (*LoadTFPlanReply, error) {
	out := new(LoadTFPlanReply)
	err := c.cc.Invoke(ctx, "/runner.Runner/LoadTFPlan", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *runnerClient) Apply(ctx context.Context, in *ApplyRequest, opts ...grpc.CallOption) (*ApplyReply, error) {
	out := new(ApplyReply)
	err := c.cc.Invoke(ctx, "/runner.Runner/Apply", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *runnerClient) GetInventory(ctx context.Context, in *GetInventoryRequest, opts ...grpc.CallOption) (*GetInventoryReply, error) {
	out := new(GetInventoryReply)
	err := c.cc.Invoke(ctx, "/runner.Runner/GetInventory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *runnerClient) Destroy(ctx context.Context, in *DestroyRequest, opts ...grpc.CallOption) (*DestroyReply, error) {
	out := new(DestroyReply)
	err := c.cc.Invoke(ctx, "/runner.Runner/Destroy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *runnerClient) Output(ctx context.Context, in *OutputRequest, opts ...grpc.CallOption) (*OutputReply, error) {
	out := new(OutputReply)
	err := c.cc.Invoke(ctx, "/runner.Runner/Output", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *runnerClient) WriteOutputs(ctx context.Context, in *WriteOutputsRequest, opts ...grpc.CallOption) (*WriteOutputsReply, error) {
	out := new(WriteOutputsReply)
	err := c.cc.Invoke(ctx, "/runner.Runner/WriteOutputs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *runnerClient) GetOutputs(ctx context.Context, in *GetOutputsRequest, opts ...grpc.CallOption) (*GetOutputsReply, error) {
	out := new(GetOutputsReply)
	err := c.cc.Invoke(ctx, "/runner.Runner/GetOutputs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *runnerClient) Init(ctx context.Context, in *InitRequest, opts ...grpc.CallOption) (*InitReply, error) {
	out := new(InitReply)
	err := c.cc.Invoke(ctx, "/runner.Runner/Init", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *runnerClient) SelectWorkspace(ctx context.Context, in *WorkspaceRequest, opts ...grpc.CallOption) (*WorkspaceReply, error) {
	out := new(WorkspaceReply)
	err := c.cc.Invoke(ctx, "/runner.Runner/SelectWorkspace", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *runnerClient) Upload(ctx context.Context, in *UploadRequest, opts ...grpc.CallOption) (*UploadReply, error) {
	out := new(UploadReply)
	err := c.cc.Invoke(ctx, "/runner.Runner/Upload", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *runnerClient) FinalizeSecrets(ctx context.Context, in *FinalizeSecretsRequest, opts ...grpc.CallOption) (*FinalizeSecretsReply, error) {
	out := new(FinalizeSecretsReply)
	err := c.cc.Invoke(ctx, "/runner.Runner/FinalizeSecrets", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *runnerClient) ForceUnlock(ctx context.Context, in *ForceUnlockRequest, opts ...grpc.CallOption) (*ForceUnlockReply, error) {
	out := new(ForceUnlockReply)
	err := c.cc.Invoke(ctx, "/runner.Runner/ForceUnlock", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RunnerServer is the server API for Runner service.
// All implementations must embed UnimplementedRunnerServer
// for forward compatibility
type RunnerServer interface {
	LookPath(context.Context, *LookPathRequest) (*LookPathReply, error)
	NewTerraform(context.Context, *NewTerraformRequest) (*NewTerraformReply, error)
	SetEnv(context.Context, *SetEnvRequest) (*SetEnvReply, error)
	CreateFileMappings(context.Context, *CreateFileMappingsRequest) (*CreateFileMappingsReply, error)
	UploadAndExtract(context.Context, *UploadAndExtractRequest) (*UploadAndExtractReply, error)
	CleanupDir(context.Context, *CleanupDirRequest) (*CleanupDirReply, error)
	WriteBackendConfig(context.Context, *WriteBackendConfigRequest) (*WriteBackendConfigReply, error)
	ProcessCliConfig(context.Context, *ProcessCliConfigRequest) (*ProcessCliConfigReply, error)
	GenerateVarsForTF(context.Context, *GenerateVarsForTFRequest) (*GenerateVarsForTFReply, error)
	Plan(context.Context, *PlanRequest) (*PlanReply, error)
	ShowPlanFileRaw(context.Context, *ShowPlanFileRawRequest) (*ShowPlanFileRawReply, error)
	ShowPlanFile(context.Context, *ShowPlanFileRequest) (*ShowPlanFileReply, error)
	SaveTFPlan(context.Context, *SaveTFPlanRequest) (*SaveTFPlanReply, error)
	LoadTFPlan(context.Context, *LoadTFPlanRequest) (*LoadTFPlanReply, error)
	Apply(context.Context, *ApplyRequest) (*ApplyReply, error)
	GetInventory(context.Context, *GetInventoryRequest) (*GetInventoryReply, error)
	Destroy(context.Context, *DestroyRequest) (*DestroyReply, error)
	Output(context.Context, *OutputRequest) (*OutputReply, error)
	WriteOutputs(context.Context, *WriteOutputsRequest) (*WriteOutputsReply, error)
	GetOutputs(context.Context, *GetOutputsRequest) (*GetOutputsReply, error)
	Init(context.Context, *InitRequest) (*InitReply, error)
	SelectWorkspace(context.Context, *WorkspaceRequest) (*WorkspaceReply, error)
	Upload(context.Context, *UploadRequest) (*UploadReply, error)
	FinalizeSecrets(context.Context, *FinalizeSecretsRequest) (*FinalizeSecretsReply, error)
	ForceUnlock(context.Context, *ForceUnlockRequest) (*ForceUnlockReply, error)
	mustEmbedUnimplementedRunnerServer()
}

// UnimplementedRunnerServer must be embedded to have forward compatible implementations.
type UnimplementedRunnerServer struct {
}

func (UnimplementedRunnerServer) LookPath(context.Context, *LookPathRequest) (*LookPathReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LookPath not implemented")
}
func (UnimplementedRunnerServer) NewTerraform(context.Context, *NewTerraformRequest) (*NewTerraformReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NewTerraform not implemented")
}
func (UnimplementedRunnerServer) SetEnv(context.Context, *SetEnvRequest) (*SetEnvReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetEnv not implemented")
}
func (UnimplementedRunnerServer) CreateFileMappings(context.Context, *CreateFileMappingsRequest) (*CreateFileMappingsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateFileMappings not implemented")
}
func (UnimplementedRunnerServer) UploadAndExtract(context.Context, *UploadAndExtractRequest) (*UploadAndExtractReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UploadAndExtract not implemented")
}
func (UnimplementedRunnerServer) CleanupDir(context.Context, *CleanupDirRequest) (*CleanupDirReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CleanupDir not implemented")
}
func (UnimplementedRunnerServer) WriteBackendConfig(context.Context, *WriteBackendConfigRequest) (*WriteBackendConfigReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WriteBackendConfig not implemented")
}
func (UnimplementedRunnerServer) ProcessCliConfig(context.Context, *ProcessCliConfigRequest) (*ProcessCliConfigReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProcessCliConfig not implemented")
}
func (UnimplementedRunnerServer) GenerateVarsForTF(context.Context, *GenerateVarsForTFRequest) (*GenerateVarsForTFReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateVarsForTF not implemented")
}
func (UnimplementedRunnerServer) Plan(context.Context, *PlanRequest) (*PlanReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Plan not implemented")
}
func (UnimplementedRunnerServer) ShowPlanFileRaw(context.Context, *ShowPlanFileRawRequest) (*ShowPlanFileRawReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ShowPlanFileRaw not implemented")
}
func (UnimplementedRunnerServer) ShowPlanFile(context.Context, *ShowPlanFileRequest) (*ShowPlanFileReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ShowPlanFile not implemented")
}
func (UnimplementedRunnerServer) SaveTFPlan(context.Context, *SaveTFPlanRequest) (*SaveTFPlanReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveTFPlan not implemented")
}
func (UnimplementedRunnerServer) LoadTFPlan(context.Context, *LoadTFPlanRequest) (*LoadTFPlanReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoadTFPlan not implemented")
}
func (UnimplementedRunnerServer) Apply(context.Context, *ApplyRequest) (*ApplyReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Apply not implemented")
}
func (UnimplementedRunnerServer) GetInventory(context.Context, *GetInventoryRequest) (*GetInventoryReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetInventory not implemented")
}
func (UnimplementedRunnerServer) Destroy(context.Context, *DestroyRequest) (*DestroyReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Destroy not implemented")
}
func (UnimplementedRunnerServer) Output(context.Context, *OutputRequest) (*OutputReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Output not implemented")
}
func (UnimplementedRunnerServer) WriteOutputs(context.Context, *WriteOutputsRequest) (*WriteOutputsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WriteOutputs not implemented")
}
func (UnimplementedRunnerServer) GetOutputs(context.Context, *GetOutputsRequest) (*GetOutputsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOutputs not implemented")
}
func (UnimplementedRunnerServer) Init(context.Context, *InitRequest) (*InitReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Init not implemented")
}
func (UnimplementedRunnerServer) SelectWorkspace(context.Context, *WorkspaceRequest) (*WorkspaceReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SelectWorkspace not implemented")
}
func (UnimplementedRunnerServer) Upload(context.Context, *UploadRequest) (*UploadReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Upload not implemented")
}
func (UnimplementedRunnerServer) FinalizeSecrets(context.Context, *FinalizeSecretsRequest) (*FinalizeSecretsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FinalizeSecrets not implemented")
}
func (UnimplementedRunnerServer) ForceUnlock(context.Context, *ForceUnlockRequest) (*ForceUnlockReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ForceUnlock not implemented")
}
func (UnimplementedRunnerServer) mustEmbedUnimplementedRunnerServer() {}

// UnsafeRunnerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RunnerServer will
// result in compilation errors.
type UnsafeRunnerServer interface {
	mustEmbedUnimplementedRunnerServer()
}

func RegisterRunnerServer(s grpc.ServiceRegistrar, srv RunnerServer) {
	s.RegisterService(&Runner_ServiceDesc, srv)
}

func _Runner_LookPath_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LookPathRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RunnerServer).LookPath(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/runner.Runner/LookPath",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RunnerServer).LookPath(ctx, req.(*LookPathRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Runner_NewTerraform_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewTerraformRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RunnerServer).NewTerraform(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/runner.Runner/NewTerraform",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RunnerServer).NewTerraform(ctx, req.(*NewTerraformRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Runner_SetEnv_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetEnvRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RunnerServer).SetEnv(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/runner.Runner/SetEnv",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RunnerServer).SetEnv(ctx, req.(*SetEnvRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Runner_CreateFileMappings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateFileMappingsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RunnerServer).CreateFileMappings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/runner.Runner/CreateFileMappings",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RunnerServer).CreateFileMappings(ctx, req.(*CreateFileMappingsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Runner_UploadAndExtract_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UploadAndExtractRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RunnerServer).UploadAndExtract(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/runner.Runner/UploadAndExtract",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RunnerServer).UploadAndExtract(ctx, req.(*UploadAndExtractRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Runner_CleanupDir_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CleanupDirRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RunnerServer).CleanupDir(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/runner.Runner/CleanupDir",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RunnerServer).CleanupDir(ctx, req.(*CleanupDirRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Runner_WriteBackendConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WriteBackendConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RunnerServer).WriteBackendConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/runner.Runner/WriteBackendConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RunnerServer).WriteBackendConfig(ctx, req.(*WriteBackendConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Runner_ProcessCliConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProcessCliConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RunnerServer).ProcessCliConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/runner.Runner/ProcessCliConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RunnerServer).ProcessCliConfig(ctx, req.(*ProcessCliConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Runner_GenerateVarsForTF_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenerateVarsForTFRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RunnerServer).GenerateVarsForTF(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/runner.Runner/GenerateVarsForTF",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RunnerServer).GenerateVarsForTF(ctx, req.(*GenerateVarsForTFRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Runner_Plan_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PlanRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RunnerServer).Plan(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/runner.Runner/Plan",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RunnerServer).Plan(ctx, req.(*PlanRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Runner_ShowPlanFileRaw_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ShowPlanFileRawRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RunnerServer).ShowPlanFileRaw(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/runner.Runner/ShowPlanFileRaw",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RunnerServer).ShowPlanFileRaw(ctx, req.(*ShowPlanFileRawRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Runner_ShowPlanFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ShowPlanFileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RunnerServer).ShowPlanFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/runner.Runner/ShowPlanFile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RunnerServer).ShowPlanFile(ctx, req.(*ShowPlanFileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Runner_SaveTFPlan_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SaveTFPlanRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RunnerServer).SaveTFPlan(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/runner.Runner/SaveTFPlan",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RunnerServer).SaveTFPlan(ctx, req.(*SaveTFPlanRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Runner_LoadTFPlan_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoadTFPlanRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RunnerServer).LoadTFPlan(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/runner.Runner/LoadTFPlan",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RunnerServer).LoadTFPlan(ctx, req.(*LoadTFPlanRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Runner_Apply_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ApplyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RunnerServer).Apply(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/runner.Runner/Apply",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RunnerServer).Apply(ctx, req.(*ApplyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Runner_GetInventory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetInventoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RunnerServer).GetInventory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/runner.Runner/GetInventory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RunnerServer).GetInventory(ctx, req.(*GetInventoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Runner_Destroy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DestroyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RunnerServer).Destroy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/runner.Runner/Destroy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RunnerServer).Destroy(ctx, req.(*DestroyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Runner_Output_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OutputRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RunnerServer).Output(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/runner.Runner/Output",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RunnerServer).Output(ctx, req.(*OutputRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Runner_WriteOutputs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WriteOutputsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RunnerServer).WriteOutputs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/runner.Runner/WriteOutputs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RunnerServer).WriteOutputs(ctx, req.(*WriteOutputsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Runner_GetOutputs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOutputsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RunnerServer).GetOutputs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/runner.Runner/GetOutputs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RunnerServer).GetOutputs(ctx, req.(*GetOutputsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Runner_Init_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InitRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RunnerServer).Init(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/runner.Runner/Init",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RunnerServer).Init(ctx, req.(*InitRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Runner_SelectWorkspace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WorkspaceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RunnerServer).SelectWorkspace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/runner.Runner/SelectWorkspace",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RunnerServer).SelectWorkspace(ctx, req.(*WorkspaceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Runner_Upload_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UploadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RunnerServer).Upload(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/runner.Runner/Upload",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RunnerServer).Upload(ctx, req.(*UploadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Runner_FinalizeSecrets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FinalizeSecretsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RunnerServer).FinalizeSecrets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/runner.Runner/FinalizeSecrets",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RunnerServer).FinalizeSecrets(ctx, req.(*FinalizeSecretsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Runner_ForceUnlock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ForceUnlockRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RunnerServer).ForceUnlock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/runner.Runner/ForceUnlock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RunnerServer).ForceUnlock(ctx, req.(*ForceUnlockRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Runner_ServiceDesc is the grpc.ServiceDesc for Runner service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Runner_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "runner.Runner",
	HandlerType: (*RunnerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "LookPath",
			Handler:    _Runner_LookPath_Handler,
		},
		{
			MethodName: "NewTerraform",
			Handler:    _Runner_NewTerraform_Handler,
		},
		{
			MethodName: "SetEnv",
			Handler:    _Runner_SetEnv_Handler,
		},
		{
			MethodName: "CreateFileMappings",
			Handler:    _Runner_CreateFileMappings_Handler,
		},
		{
			MethodName: "UploadAndExtract",
			Handler:    _Runner_UploadAndExtract_Handler,
		},
		{
			MethodName: "CleanupDir",
			Handler:    _Runner_CleanupDir_Handler,
		},
		{
			MethodName: "WriteBackendConfig",
			Handler:    _Runner_WriteBackendConfig_Handler,
		},
		{
			MethodName: "ProcessCliConfig",
			Handler:    _Runner_ProcessCliConfig_Handler,
		},
		{
			MethodName: "GenerateVarsForTF",
			Handler:    _Runner_GenerateVarsForTF_Handler,
		},
		{
			MethodName: "Plan",
			Handler:    _Runner_Plan_Handler,
		},
		{
			MethodName: "ShowPlanFileRaw",
			Handler:    _Runner_ShowPlanFileRaw_Handler,
		},
		{
			MethodName: "ShowPlanFile",
			Handler:    _Runner_ShowPlanFile_Handler,
		},
		{
			MethodName: "SaveTFPlan",
			Handler:    _Runner_SaveTFPlan_Handler,
		},
		{
			MethodName: "LoadTFPlan",
			Handler:    _Runner_LoadTFPlan_Handler,
		},
		{
			MethodName: "Apply",
			Handler:    _Runner_Apply_Handler,
		},
		{
			MethodName: "GetInventory",
			Handler:    _Runner_GetInventory_Handler,
		},
		{
			MethodName: "Destroy",
			Handler:    _Runner_Destroy_Handler,
		},
		{
			MethodName: "Output",
			Handler:    _Runner_Output_Handler,
		},
		{
			MethodName: "WriteOutputs",
			Handler:    _Runner_WriteOutputs_Handler,
		},
		{
			MethodName: "GetOutputs",
			Handler:    _Runner_GetOutputs_Handler,
		},
		{
			MethodName: "Init",
			Handler:    _Runner_Init_Handler,
		},
		{
			MethodName: "SelectWorkspace",
			Handler:    _Runner_SelectWorkspace_Handler,
		},
		{
			MethodName: "Upload",
			Handler:    _Runner_Upload_Handler,
		},
		{
			MethodName: "FinalizeSecrets",
			Handler:    _Runner_FinalizeSecrets_Handler,
		},
		{
			MethodName: "ForceUnlock",
			Handler:    _Runner_ForceUnlock_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "runner/runner.proto",
}
