# Enumerations

**NB** All enumerations also have an extra values ending with `_Force32` that is here to ensure that all compilers use the same underlying representation for the enum values. This value must never be used.


## AdapterType { #WGPUAdapterType }


TODO

```C
enum WGPUAdapterType {
        
TODO

    WGPUAdapterType_DiscreteGPU = 0x00000000,
        
TODO

    WGPUAdapterType_IntegratedGPU = 0x00000001,
        
TODO

    WGPUAdapterType_CPU = 0x00000002,
        
TODO


    WGPUAdapterType_Unknown = 0x00000003,
}
```


## AddressMode { #WGPUAddressMode }


TODO

```C
enum WGPUAddressMode {
        
TODO

    WGPUAddressMode_Repeat = 0x00000000,
        
TODO

    WGPUAddressMode_MirrorRepeat = 0x00000001,
        
TODO


    WGPUAddressMode_ClampToEdge = 0x00000002,
}
```


## BackendType { #WGPUBackendType }


TODO

```C
enum WGPUBackendType {
        
TODO

    WGPUBackendType_Undefined = 0x00000000,
        
TODO

    WGPUBackendType_Null = 0x00000001,
        
TODO

    WGPUBackendType_WebGPU = 0x00000002,
        
TODO

    WGPUBackendType_D3D11 = 0x00000003,
        
TODO

    WGPUBackendType_D3D12 = 0x00000004,
        
TODO

    WGPUBackendType_Metal = 0x00000005,
        
TODO

    WGPUBackendType_Vulkan = 0x00000006,
        
TODO

    WGPUBackendType_OpenGL = 0x00000007,
        
TODO


    WGPUBackendType_OpenGLES = 0x00000008,
}
```


## BlendFactor { #WGPUBlendFactor }


TODO

```C
enum WGPUBlendFactor {
        
TODO

    WGPUBlendFactor_Zero = 0x00000000,
        
TODO

    WGPUBlendFactor_One = 0x00000001,
        
TODO

    WGPUBlendFactor_Src = 0x00000002,
        
TODO

    WGPUBlendFactor_OneMinusSrc = 0x00000003,
        
TODO

    WGPUBlendFactor_SrcAlpha = 0x00000004,
        
TODO

    WGPUBlendFactor_OneMinusSrcAlpha = 0x00000005,
        
TODO

    WGPUBlendFactor_Dst = 0x00000006,
        
TODO

    WGPUBlendFactor_OneMinusDst = 0x00000007,
        
TODO

    WGPUBlendFactor_DstAlpha = 0x00000008,
        
TODO

    WGPUBlendFactor_OneMinusDstAlpha = 0x00000009,
        
TODO

    WGPUBlendFactor_SrcAlphaSaturated = 0x0000000A,
        
TODO

    WGPUBlendFactor_Constant = 0x0000000B,
        
TODO


    WGPUBlendFactor_OneMinusConstant = 0x0000000C,
}
```


## BlendOperation { #WGPUBlendOperation }


TODO

```C
enum WGPUBlendOperation {
        
TODO

    WGPUBlendOperation_Add = 0x00000000,
        
TODO

    WGPUBlendOperation_Subtract = 0x00000001,
        
TODO

    WGPUBlendOperation_ReverseSubtract = 0x00000002,
        
TODO

    WGPUBlendOperation_Min = 0x00000003,
        
TODO


    WGPUBlendOperation_Max = 0x00000004,
}
```


## BufferBindingType { #WGPUBufferBindingType }


TODO

```C
enum WGPUBufferBindingType {
        
TODO

    WGPUBufferBindingType_Undefined = 0x00000000,
        
TODO

    WGPUBufferBindingType_Uniform = 0x00000001,
        
TODO

    WGPUBufferBindingType_Storage = 0x00000002,
        
TODO


    WGPUBufferBindingType_ReadOnlyStorage = 0x00000003,
}
```


## BufferMapAsyncStatus { #WGPUBufferMapAsyncStatus }


TODO

```C
enum WGPUBufferMapAsyncStatus {
        
TODO

    WGPUBufferMapAsyncStatus_Success = 0x00000000,
        
TODO

    WGPUBufferMapAsyncStatus_ValidationError = 0x00000001,
        
TODO

    WGPUBufferMapAsyncStatus_Unknown = 0x00000002,
        
TODO

    WGPUBufferMapAsyncStatus_DeviceLost = 0x00000003,
        
TODO

    WGPUBufferMapAsyncStatus_DestroyedBeforeCallback = 0x00000004,
        
TODO

    WGPUBufferMapAsyncStatus_UnmappedBeforeCallback = 0x00000005,
        
TODO

    WGPUBufferMapAsyncStatus_MappingAlreadyPending = 0x00000006,
        
TODO

    WGPUBufferMapAsyncStatus_OffsetOutOfRange = 0x00000007,
        
TODO


    WGPUBufferMapAsyncStatus_SizeOutOfRange = 0x00000008,
}
```


## BufferMapState { #WGPUBufferMapState }


TODO

```C
enum WGPUBufferMapState {
        
TODO

    WGPUBufferMapState_Unmapped = 0x00000000,
        
TODO

    WGPUBufferMapState_Pending = 0x00000001,
        
TODO


    WGPUBufferMapState_Mapped = 0x00000002,
}
```


## CompareFunction { #WGPUCompareFunction }


TODO

```C
enum WGPUCompareFunction {
        
TODO

    WGPUCompareFunction_Undefined = 0x00000000,
        
TODO

    WGPUCompareFunction_Never = 0x00000001,
        
TODO

    WGPUCompareFunction_Less = 0x00000002,
        
TODO

    WGPUCompareFunction_LessEqual = 0x00000003,
        
TODO

    WGPUCompareFunction_Greater = 0x00000004,
        
TODO

    WGPUCompareFunction_GreaterEqual = 0x00000005,
        
TODO

    WGPUCompareFunction_Equal = 0x00000006,
        
TODO

    WGPUCompareFunction_NotEqual = 0x00000007,
        
TODO


    WGPUCompareFunction_Always = 0x00000008,
}
```


## CompilationInfoRequestStatus { #WGPUCompilationInfoRequestStatus }


TODO

```C
enum WGPUCompilationInfoRequestStatus {
        
TODO

    WGPUCompilationInfoRequestStatus_Success = 0x00000000,
        
TODO

    WGPUCompilationInfoRequestStatus_Error = 0x00000001,
        
TODO

    WGPUCompilationInfoRequestStatus_DeviceLost = 0x00000002,
        
TODO


    WGPUCompilationInfoRequestStatus_Unknown = 0x00000003,
}
```


## CompilationMessageType { #WGPUCompilationMessageType }


TODO

```C
enum WGPUCompilationMessageType {
        
TODO

    WGPUCompilationMessageType_Error = 0x00000000,
        
TODO

    WGPUCompilationMessageType_Warning = 0x00000001,
        
TODO


    WGPUCompilationMessageType_Info = 0x00000002,
}
```


## CompositeAlphaMode { #WGPUCompositeAlphaMode }


TODO

```C
enum WGPUCompositeAlphaMode {
        
TODO

    WGPUCompositeAlphaMode_Auto = 0x00000000,
        
TODO

    WGPUCompositeAlphaMode_Opaque = 0x00000001,
        
TODO

    WGPUCompositeAlphaMode_Premultiplied = 0x00000002,
        
TODO

    WGPUCompositeAlphaMode_Unpremultiplied = 0x00000003,
        
TODO


    WGPUCompositeAlphaMode_Inherit = 0x00000004,
}
```


## CreatePipelineAsyncStatus { #WGPUCreatePipelineAsyncStatus }


TODO

```C
enum WGPUCreatePipelineAsyncStatus {
        
TODO

    WGPUCreatePipelineAsyncStatus_Success = 0x00000000,
        
TODO

    WGPUCreatePipelineAsyncStatus_ValidationError = 0x00000001,
        
TODO

    WGPUCreatePipelineAsyncStatus_InternalError = 0x00000002,
        
TODO

    WGPUCreatePipelineAsyncStatus_DeviceLost = 0x00000003,
        
TODO

    WGPUCreatePipelineAsyncStatus_DeviceDestroyed = 0x00000004,
        
TODO


    WGPUCreatePipelineAsyncStatus_Unknown = 0x00000005,
}
```


## CullMode { #WGPUCullMode }


TODO

```C
enum WGPUCullMode {
        
TODO

    WGPUCullMode_None = 0x00000000,
        
TODO

    WGPUCullMode_Front = 0x00000001,
        
TODO


    WGPUCullMode_Back = 0x00000002,
}
```


## DeviceLostReason { #WGPUDeviceLostReason }


TODO

```C
enum WGPUDeviceLostReason {
        
TODO

    WGPUDeviceLostReason_Undefined = 0x00000000,
        
TODO


    WGPUDeviceLostReason_Destroyed = 0x00000001,
}
```


## ErrorFilter { #WGPUErrorFilter }


TODO

```C
enum WGPUErrorFilter {
        
TODO

    WGPUErrorFilter_Validation = 0x00000000,
        
TODO

    WGPUErrorFilter_OutOfMemory = 0x00000001,
        
TODO


    WGPUErrorFilter_Internal = 0x00000002,
}
```


## ErrorType { #WGPUErrorType }


TODO

```C
enum WGPUErrorType {
        
TODO

    WGPUErrorType_NoError = 0x00000000,
        
TODO

    WGPUErrorType_Validation = 0x00000001,
        
TODO

    WGPUErrorType_OutOfMemory = 0x00000002,
        
TODO

    WGPUErrorType_Internal = 0x00000003,
        
TODO

    WGPUErrorType_Unknown = 0x00000004,
        
TODO


    WGPUErrorType_DeviceLost = 0x00000005,
}
```


## FeatureName { #WGPUFeatureName }


TODO

```C
enum WGPUFeatureName {
        
TODO

    WGPUFeatureName_Undefined = 0x00000000,
        
TODO

    WGPUFeatureName_DepthClipControl = 0x00000001,
        
TODO

    WGPUFeatureName_Depth32FloatStencil8 = 0x00000002,
        
TODO

    WGPUFeatureName_TimestampQuery = 0x00000003,
        
TODO

    WGPUFeatureName_TextureCompressionBC = 0x00000004,
        
TODO

    WGPUFeatureName_TextureCompressionETC2 = 0x00000005,
        
TODO

    WGPUFeatureName_TextureCompressionASTC = 0x00000006,
        
TODO

    WGPUFeatureName_IndirectFirstInstance = 0x00000007,
        
TODO

    WGPUFeatureName_ShaderF16 = 0x00000008,
        
TODO

    WGPUFeatureName_RG11B10UfloatRenderable = 0x00000009,
        
TODO

    WGPUFeatureName_BGRA8UnormStorage = 0x0000000A,
        
TODO


    WGPUFeatureName_Float32Filterable = 0x0000000B,
}
```


## FilterMode { #WGPUFilterMode }


TODO

```C
enum WGPUFilterMode {
        
TODO

    WGPUFilterMode_Nearest = 0x00000000,
        
TODO


    WGPUFilterMode_Linear = 0x00000001,
}
```


## FrontFace { #WGPUFrontFace }


TODO

```C
enum WGPUFrontFace {
        
TODO

    WGPUFrontFace_CCW = 0x00000000,
        
TODO


    WGPUFrontFace_CW = 0x00000001,
}
```


## IndexFormat { #WGPUIndexFormat }


TODO

```C
enum WGPUIndexFormat {
        
TODO

    WGPUIndexFormat_Undefined = 0x00000000,
        
TODO

    WGPUIndexFormat_Uint16 = 0x00000001,
        
TODO


    WGPUIndexFormat_Uint32 = 0x00000002,
}
```


## LoadOp { #WGPULoadOp }


TODO

```C
enum WGPULoadOp {
        
TODO

    WGPULoadOp_Undefined = 0x00000000,
        
TODO

    WGPULoadOp_Clear = 0x00000001,
        
TODO


    WGPULoadOp_Load = 0x00000002,
}
```


## MipmapFilterMode { #WGPUMipmapFilterMode }


TODO

```C
enum WGPUMipmapFilterMode {
        
TODO

    WGPUMipmapFilterMode_Nearest = 0x00000000,
        
TODO


    WGPUMipmapFilterMode_Linear = 0x00000001,
}
```


## PowerPreference { #WGPUPowerPreference }


TODO

```C
enum WGPUPowerPreference {
        
TODO

    WGPUPowerPreference_Undefined = 0x00000000,
        
TODO

    WGPUPowerPreference_LowPower = 0x00000001,
        
TODO


    WGPUPowerPreference_HighPerformance = 0x00000002,
}
```


## PresentMode { #WGPUPresentMode }


TODO

```C
enum WGPUPresentMode {
        
TODO

    WGPUPresentMode_Fifo = 0x00000000,
        
TODO

    WGPUPresentMode_FifoRelaxed = 0x00000001,
        
TODO

    WGPUPresentMode_Immediate = 0x00000002,
        
TODO


    WGPUPresentMode_Mailbox = 0x00000003,
}
```


## PrimitiveTopology { #WGPUPrimitiveTopology }


TODO

```C
enum WGPUPrimitiveTopology {
        
TODO

    WGPUPrimitiveTopology_PointList = 0x00000000,
        
TODO

    WGPUPrimitiveTopology_LineList = 0x00000001,
        
TODO

    WGPUPrimitiveTopology_LineStrip = 0x00000002,
        
TODO

    WGPUPrimitiveTopology_TriangleList = 0x00000003,
        
TODO


    WGPUPrimitiveTopology_TriangleStrip = 0x00000004,
}
```


## QueryType { #WGPUQueryType }


TODO

```C
enum WGPUQueryType {
        
TODO

    WGPUQueryType_Occlusion = 0x00000000,
        
TODO


    WGPUQueryType_Timestamp = 0x00000001,
}
```


## QueueWorkDoneStatus { #WGPUQueueWorkDoneStatus }


TODO

```C
enum WGPUQueueWorkDoneStatus {
        
TODO

    WGPUQueueWorkDoneStatus_Success = 0x00000000,
        
TODO

    WGPUQueueWorkDoneStatus_Error = 0x00000001,
        
TODO

    WGPUQueueWorkDoneStatus_Unknown = 0x00000002,
        
TODO


    WGPUQueueWorkDoneStatus_DeviceLost = 0x00000003,
}
```


## RequestAdapterStatus { #WGPURequestAdapterStatus }


TODO

```C
enum WGPURequestAdapterStatus {
        
TODO

    WGPURequestAdapterStatus_Success = 0x00000000,
        
TODO

    WGPURequestAdapterStatus_Unavailable = 0x00000001,
        
TODO

    WGPURequestAdapterStatus_Error = 0x00000002,
        
TODO


    WGPURequestAdapterStatus_Unknown = 0x00000003,
}
```


## RequestDeviceStatus { #WGPURequestDeviceStatus }


TODO

```C
enum WGPURequestDeviceStatus {
        
TODO

    WGPURequestDeviceStatus_Success = 0x00000000,
        
TODO

    WGPURequestDeviceStatus_Error = 0x00000001,
        
TODO


    WGPURequestDeviceStatus_Unknown = 0x00000002,
}
```


## SType { #WGPUSType }


TODO

```C
enum WGPUSType {
        
TODO

    WGPUSType_Invalid = 0x00000000,
        
TODO

    WGPUSType_SurfaceDescriptorFromMetalLayer = 0x00000001,
        
TODO

    WGPUSType_SurfaceDescriptorFromWindowsHWND = 0x00000002,
        
TODO

    WGPUSType_SurfaceDescriptorFromXlibWindow = 0x00000003,
        
TODO

    WGPUSType_SurfaceDescriptorFromCanvasHTMLSelector = 0x00000004,
        
TODO

    WGPUSType_ShaderModuleSPIRVDescriptor = 0x00000005,
        
TODO

    WGPUSType_ShaderModuleWGSLDescriptor = 0x00000006,
        
TODO

    WGPUSType_PrimitiveDepthClipControl = 0x00000007,
        
TODO

    WGPUSType_SurfaceDescriptorFromWaylandSurface = 0x00000008,
        
TODO

    WGPUSType_SurfaceDescriptorFromAndroidNativeWindow = 0x00000009,
        
TODO

    WGPUSType_SurfaceDescriptorFromXcbWindow = 0x0000000A,
        
TODO


    WGPUSType_RenderPassDescriptorMaxDrawCount = 0x0000000F,
}
```


## SamplerBindingType { #WGPUSamplerBindingType }


TODO

```C
enum WGPUSamplerBindingType {
        
TODO

    WGPUSamplerBindingType_Undefined = 0x00000000,
        
TODO

    WGPUSamplerBindingType_Filtering = 0x00000001,
        
TODO

    WGPUSamplerBindingType_NonFiltering = 0x00000002,
        
TODO


    WGPUSamplerBindingType_Comparison = 0x00000003,
}
```


## StencilOperation { #WGPUStencilOperation }


TODO

```C
enum WGPUStencilOperation {
        
TODO

    WGPUStencilOperation_Keep = 0x00000000,
        
TODO

    WGPUStencilOperation_Zero = 0x00000001,
        
TODO

    WGPUStencilOperation_Replace = 0x00000002,
        
TODO

    WGPUStencilOperation_Invert = 0x00000003,
        
TODO

    WGPUStencilOperation_IncrementClamp = 0x00000004,
        
TODO

    WGPUStencilOperation_DecrementClamp = 0x00000005,
        
TODO

    WGPUStencilOperation_IncrementWrap = 0x00000006,
        
TODO


    WGPUStencilOperation_DecrementWrap = 0x00000007,
}
```


## StorageTextureAccess { #WGPUStorageTextureAccess }


TODO

```C
enum WGPUStorageTextureAccess {
        
TODO

    WGPUStorageTextureAccess_Undefined = 0x00000000,
        
TODO

    WGPUStorageTextureAccess_WriteOnly = 0x00000001,
        
TODO

    WGPUStorageTextureAccess_ReadOnly = 0x00000002,
        
TODO


    WGPUStorageTextureAccess_ReadWrite = 0x00000003,
}
```


## StoreOp { #WGPUStoreOp }


TODO

```C
enum WGPUStoreOp {
        
TODO

    WGPUStoreOp_Undefined = 0x00000000,
        
TODO

    WGPUStoreOp_Store = 0x00000001,
        
TODO


    WGPUStoreOp_Discard = 0x00000002,
}
```


## SurfaceGetCurrentTextureStatus { #WGPUSurfaceGetCurrentTextureStatus }


TODO

```C
enum WGPUSurfaceGetCurrentTextureStatus {
        
TODO

    WGPUSurfaceGetCurrentTextureStatus_Success = 0x00000000,
        
TODO

    WGPUSurfaceGetCurrentTextureStatus_Timeout = 0x00000001,
        
TODO

    WGPUSurfaceGetCurrentTextureStatus_Outdated = 0x00000002,
        
TODO

    WGPUSurfaceGetCurrentTextureStatus_Lost = 0x00000003,
        
TODO

    WGPUSurfaceGetCurrentTextureStatus_OutOfMemory = 0x00000004,
        
TODO


    WGPUSurfaceGetCurrentTextureStatus_DeviceLost = 0x00000005,
}
```


## TextureAspect { #WGPUTextureAspect }


TODO

```C
enum WGPUTextureAspect {
        
TODO

    WGPUTextureAspect_All = 0x00000000,
        
TODO

    WGPUTextureAspect_StencilOnly = 0x00000001,
        
TODO


    WGPUTextureAspect_DepthOnly = 0x00000002,
}
```


## TextureDimension { #WGPUTextureDimension }


TODO

```C
enum WGPUTextureDimension {
        
TODO

    WGPUTextureDimension_1D = 0x00000000,
        
TODO

    WGPUTextureDimension_2D = 0x00000001,
        
TODO


    WGPUTextureDimension_3D = 0x00000002,
}
```


## TextureFormat { #WGPUTextureFormat }


TODO

```C
enum WGPUTextureFormat {
        
TODO

    WGPUTextureFormat_Undefined = 0x00000000,
        
TODO

    WGPUTextureFormat_R8Unorm = 0x00000001,
        
TODO

    WGPUTextureFormat_R8Snorm = 0x00000002,
        
TODO

    WGPUTextureFormat_R8Uint = 0x00000003,
        
TODO

    WGPUTextureFormat_R8Sint = 0x00000004,
        
TODO

    WGPUTextureFormat_R16Uint = 0x00000005,
        
TODO

    WGPUTextureFormat_R16Sint = 0x00000006,
        
TODO

    WGPUTextureFormat_R16Float = 0x00000007,
        
TODO

    WGPUTextureFormat_RG8Unorm = 0x00000008,
        
TODO

    WGPUTextureFormat_RG8Snorm = 0x00000009,
        
TODO

    WGPUTextureFormat_RG8Uint = 0x0000000A,
        
TODO

    WGPUTextureFormat_RG8Sint = 0x0000000B,
        
TODO

    WGPUTextureFormat_R32Float = 0x0000000C,
        
TODO

    WGPUTextureFormat_R32Uint = 0x0000000D,
        
TODO

    WGPUTextureFormat_R32Sint = 0x0000000E,
        
TODO

    WGPUTextureFormat_RG16Uint = 0x0000000F,
        
TODO

    WGPUTextureFormat_RG16Sint = 0x00000010,
        
TODO

    WGPUTextureFormat_RG16Float = 0x00000011,
        
TODO

    WGPUTextureFormat_RGBA8Unorm = 0x00000012,
        
TODO

    WGPUTextureFormat_RGBA8UnormSrgb = 0x00000013,
        
TODO

    WGPUTextureFormat_RGBA8Snorm = 0x00000014,
        
TODO

    WGPUTextureFormat_RGBA8Uint = 0x00000015,
        
TODO

    WGPUTextureFormat_RGBA8Sint = 0x00000016,
        
TODO

    WGPUTextureFormat_BGRA8Unorm = 0x00000017,
        
TODO

    WGPUTextureFormat_BGRA8UnormSrgb = 0x00000018,
        
TODO

    WGPUTextureFormat_RGB10A2Uint = 0x00000019,
        
TODO

    WGPUTextureFormat_RGB10A2Unorm = 0x0000001A,
        
TODO

    WGPUTextureFormat_RG11B10Ufloat = 0x0000001B,
        
TODO

    WGPUTextureFormat_RGB9E5Ufloat = 0x0000001C,
        
TODO

    WGPUTextureFormat_RG32Float = 0x0000001D,
        
TODO

    WGPUTextureFormat_RG32Uint = 0x0000001E,
        
TODO

    WGPUTextureFormat_RG32Sint = 0x0000001F,
        
TODO

    WGPUTextureFormat_RGBA16Uint = 0x00000020,
        
TODO

    WGPUTextureFormat_RGBA16Sint = 0x00000021,
        
TODO

    WGPUTextureFormat_RGBA16Float = 0x00000022,
        
TODO

    WGPUTextureFormat_RGBA32Float = 0x00000023,
        
TODO

    WGPUTextureFormat_RGBA32Uint = 0x00000024,
        
TODO

    WGPUTextureFormat_RGBA32Sint = 0x00000025,
        
TODO

    WGPUTextureFormat_Stencil8 = 0x00000026,
        
TODO

    WGPUTextureFormat_Depth16Unorm = 0x00000027,
        
TODO

    WGPUTextureFormat_Depth24Plus = 0x00000028,
        
TODO

    WGPUTextureFormat_Depth24PlusStencil8 = 0x00000029,
        
TODO

    WGPUTextureFormat_Depth32Float = 0x0000002A,
        
TODO

    WGPUTextureFormat_Depth32FloatStencil8 = 0x0000002B,
        
TODO

    WGPUTextureFormat_BC1RGBAUnorm = 0x0000002C,
        
TODO

    WGPUTextureFormat_BC1RGBAUnormSrgb = 0x0000002D,
        
TODO

    WGPUTextureFormat_BC2RGBAUnorm = 0x0000002E,
        
TODO

    WGPUTextureFormat_BC2RGBAUnormSrgb = 0x0000002F,
        
TODO

    WGPUTextureFormat_BC3RGBAUnorm = 0x00000030,
        
TODO

    WGPUTextureFormat_BC3RGBAUnormSrgb = 0x00000031,
        
TODO

    WGPUTextureFormat_BC4RUnorm = 0x00000032,
        
TODO

    WGPUTextureFormat_BC4RSnorm = 0x00000033,
        
TODO

    WGPUTextureFormat_BC5RGUnorm = 0x00000034,
        
TODO

    WGPUTextureFormat_BC5RGSnorm = 0x00000035,
        
TODO

    WGPUTextureFormat_BC6HRGBUfloat = 0x00000036,
        
TODO

    WGPUTextureFormat_BC6HRGBFloat = 0x00000037,
        
TODO

    WGPUTextureFormat_BC7RGBAUnorm = 0x00000038,
        
TODO

    WGPUTextureFormat_BC7RGBAUnormSrgb = 0x00000039,
        
TODO

    WGPUTextureFormat_ETC2RGB8Unorm = 0x0000003A,
        
TODO

    WGPUTextureFormat_ETC2RGB8UnormSrgb = 0x0000003B,
        
TODO

    WGPUTextureFormat_ETC2RGB8A1Unorm = 0x0000003C,
        
TODO

    WGPUTextureFormat_ETC2RGB8A1UnormSrgb = 0x0000003D,
        
TODO

    WGPUTextureFormat_ETC2RGBA8Unorm = 0x0000003E,
        
TODO

    WGPUTextureFormat_ETC2RGBA8UnormSrgb = 0x0000003F,
        
TODO

    WGPUTextureFormat_EACR11Unorm = 0x00000040,
        
TODO

    WGPUTextureFormat_EACR11Snorm = 0x00000041,
        
TODO

    WGPUTextureFormat_EACRG11Unorm = 0x00000042,
        
TODO

    WGPUTextureFormat_EACRG11Snorm = 0x00000043,
        
TODO

    WGPUTextureFormat_ASTC4x4Unorm = 0x00000044,
        
TODO

    WGPUTextureFormat_ASTC4x4UnormSrgb = 0x00000045,
        
TODO

    WGPUTextureFormat_ASTC5x4Unorm = 0x00000046,
        
TODO

    WGPUTextureFormat_ASTC5x4UnormSrgb = 0x00000047,
        
TODO

    WGPUTextureFormat_ASTC5x5Unorm = 0x00000048,
        
TODO

    WGPUTextureFormat_ASTC5x5UnormSrgb = 0x00000049,
        
TODO

    WGPUTextureFormat_ASTC6x5Unorm = 0x0000004A,
        
TODO

    WGPUTextureFormat_ASTC6x5UnormSrgb = 0x0000004B,
        
TODO

    WGPUTextureFormat_ASTC6x6Unorm = 0x0000004C,
        
TODO

    WGPUTextureFormat_ASTC6x6UnormSrgb = 0x0000004D,
        
TODO

    WGPUTextureFormat_ASTC8x5Unorm = 0x0000004E,
        
TODO

    WGPUTextureFormat_ASTC8x5UnormSrgb = 0x0000004F,
        
TODO

    WGPUTextureFormat_ASTC8x6Unorm = 0x00000050,
        
TODO

    WGPUTextureFormat_ASTC8x6UnormSrgb = 0x00000051,
        
TODO

    WGPUTextureFormat_ASTC8x8Unorm = 0x00000052,
        
TODO

    WGPUTextureFormat_ASTC8x8UnormSrgb = 0x00000053,
        
TODO

    WGPUTextureFormat_ASTC10x5Unorm = 0x00000054,
        
TODO

    WGPUTextureFormat_ASTC10x5UnormSrgb = 0x00000055,
        
TODO

    WGPUTextureFormat_ASTC10x6Unorm = 0x00000056,
        
TODO

    WGPUTextureFormat_ASTC10x6UnormSrgb = 0x00000057,
        
TODO

    WGPUTextureFormat_ASTC10x8Unorm = 0x00000058,
        
TODO

    WGPUTextureFormat_ASTC10x8UnormSrgb = 0x00000059,
        
TODO

    WGPUTextureFormat_ASTC10x10Unorm = 0x0000005A,
        
TODO

    WGPUTextureFormat_ASTC10x10UnormSrgb = 0x0000005B,
        
TODO

    WGPUTextureFormat_ASTC12x10Unorm = 0x0000005C,
        
TODO

    WGPUTextureFormat_ASTC12x10UnormSrgb = 0x0000005D,
        
TODO

    WGPUTextureFormat_ASTC12x12Unorm = 0x0000005E,
        
TODO


    WGPUTextureFormat_ASTC12x12UnormSrgb = 0x0000005F,
}
```


## TextureSampleType { #WGPUTextureSampleType }


TODO

```C
enum WGPUTextureSampleType {
        
TODO

    WGPUTextureSampleType_Undefined = 0x00000000,
        
TODO

    WGPUTextureSampleType_Float = 0x00000001,
        
TODO

    WGPUTextureSampleType_UnfilterableFloat = 0x00000002,
        
TODO

    WGPUTextureSampleType_Depth = 0x00000003,
        
TODO

    WGPUTextureSampleType_Sint = 0x00000004,
        
TODO


    WGPUTextureSampleType_Uint = 0x00000005,
}
```


## TextureViewDimension { #WGPUTextureViewDimension }


TODO

```C
enum WGPUTextureViewDimension {
        
TODO

    WGPUTextureViewDimension_Undefined = 0x00000000,
        
TODO

    WGPUTextureViewDimension_1D = 0x00000001,
        
TODO

    WGPUTextureViewDimension_2D = 0x00000002,
        
TODO

    WGPUTextureViewDimension_2DArray = 0x00000003,
        
TODO

    WGPUTextureViewDimension_Cube = 0x00000004,
        
TODO

    WGPUTextureViewDimension_CubeArray = 0x00000005,
        
TODO


    WGPUTextureViewDimension_3D = 0x00000006,
}
```


## VertexFormat { #WGPUVertexFormat }


TODO

```C
enum WGPUVertexFormat {
        
TODO

    WGPUVertexFormat_Undefined = 0x00000000,
        
TODO

    WGPUVertexFormat_Uint8x2 = 0x00000001,
        
TODO

    WGPUVertexFormat_Uint8x4 = 0x00000002,
        
TODO

    WGPUVertexFormat_Sint8x2 = 0x00000003,
        
TODO

    WGPUVertexFormat_Sint8x4 = 0x00000004,
        
TODO

    WGPUVertexFormat_Unorm8x2 = 0x00000005,
        
TODO

    WGPUVertexFormat_Unorm8x4 = 0x00000006,
        
TODO

    WGPUVertexFormat_Snorm8x2 = 0x00000007,
        
TODO

    WGPUVertexFormat_Snorm8x4 = 0x00000008,
        
TODO

    WGPUVertexFormat_Uint16x2 = 0x00000009,
        
TODO

    WGPUVertexFormat_Uint16x4 = 0x0000000A,
        
TODO

    WGPUVertexFormat_Sint16x2 = 0x0000000B,
        
TODO

    WGPUVertexFormat_Sint16x4 = 0x0000000C,
        
TODO

    WGPUVertexFormat_Unorm16x2 = 0x0000000D,
        
TODO

    WGPUVertexFormat_Unorm16x4 = 0x0000000E,
        
TODO

    WGPUVertexFormat_Snorm16x2 = 0x0000000F,
        
TODO

    WGPUVertexFormat_Snorm16x4 = 0x00000010,
        
TODO

    WGPUVertexFormat_Float16x2 = 0x00000011,
        
TODO

    WGPUVertexFormat_Float16x4 = 0x00000012,
        
TODO

    WGPUVertexFormat_Float32 = 0x00000013,
        
TODO

    WGPUVertexFormat_Float32x2 = 0x00000014,
        
TODO

    WGPUVertexFormat_Float32x3 = 0x00000015,
        
TODO

    WGPUVertexFormat_Float32x4 = 0x00000016,
        
TODO

    WGPUVertexFormat_Uint32 = 0x00000017,
        
TODO

    WGPUVertexFormat_Uint32x2 = 0x00000018,
        
TODO

    WGPUVertexFormat_Uint32x3 = 0x00000019,
        
TODO

    WGPUVertexFormat_Uint32x4 = 0x0000001A,
        
TODO

    WGPUVertexFormat_Sint32 = 0x0000001B,
        
TODO

    WGPUVertexFormat_Sint32x2 = 0x0000001C,
        
TODO

    WGPUVertexFormat_Sint32x3 = 0x0000001D,
        
TODO


    WGPUVertexFormat_Sint32x4 = 0x0000001E,
}
```


## VertexStepMode { #WGPUVertexStepMode }


TODO

```C
enum WGPUVertexStepMode {
        
TODO

    WGPUVertexStepMode_Vertex = 0x00000000,
        
TODO

    WGPUVertexStepMode_Instance = 0x00000001,
        
TODO


    WGPUVertexStepMode_VertexBufferNotUsed = 0x00000002,
}
```


## WGSLFeatureName { #WGPUWGSLFeatureName }


TODO

```C
enum WGPUWGSLFeatureName {
        
TODO

    WGPUWGSLFeatureName_Undefined = 0x00000000,
        
TODO

    WGPUWGSLFeatureName_ReadonlyAndReadwriteStorageTextures = 0x00000001,
        
TODO

    WGPUWGSLFeatureName_Packed4x8IntegerDotProduct = 0x00000002,
        
TODO

    WGPUWGSLFeatureName_UnrestrictedPointerParameters = 0x00000003,
        
TODO


    WGPUWGSLFeatureName_PointerCompositeAccess = 0x00000004,
}
```


