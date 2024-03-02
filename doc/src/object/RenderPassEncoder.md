

# RenderPassEncoder { #WGPURenderPassEncoder }


TODO




<br/><!-- poor man's styling, just for the demo before we use a non default theme -->
***

## BeginOcclusionQuery { #wgpuRenderPassEncoderBeginOcclusionQuery }

```C
void wgpuRenderPassEncoderBeginOcclusionQuery(WGPURenderPassEncoder renderPassEncoder, uint32_t queryIndex)
```


TODO


**Arguments:**

 - TODO




<br/><!-- poor man's styling, just for the demo before we use a non default theme -->
***

## Draw { #wgpuRenderPassEncoderDraw }

```C
void wgpuRenderPassEncoderDraw(WGPURenderPassEncoder renderPassEncoder, uint32_t vertexCount, uint32_t instanceCount, uint32_t firstVertex, uint32_t firstInstance)
```


TODO


**Arguments:**

 - TODO




<br/><!-- poor man's styling, just for the demo before we use a non default theme -->
***

## DrawIndexed { #wgpuRenderPassEncoderDrawIndexed }

```C
void wgpuRenderPassEncoderDrawIndexed(WGPURenderPassEncoder renderPassEncoder, uint32_t indexCount, uint32_t instanceCount, uint32_t firstIndex, int32_t baseVertex, uint32_t firstInstance)
```


TODO


**Arguments:**

 - TODO




<br/><!-- poor man's styling, just for the demo before we use a non default theme -->
***

## DrawIndexedIndirect { #wgpuRenderPassEncoderDrawIndexedIndirect }

```C
void wgpuRenderPassEncoderDrawIndexedIndirect(WGPURenderPassEncoder renderPassEncoder, WGPUBuffer indirectBuffer, uint64_t indirectOffset)
```


TODO


**Arguments:**

 - TODO




<br/><!-- poor man's styling, just for the demo before we use a non default theme -->
***

## DrawIndirect { #wgpuRenderPassEncoderDrawIndirect }

```C
void wgpuRenderPassEncoderDrawIndirect(WGPURenderPassEncoder renderPassEncoder, WGPUBuffer indirectBuffer, uint64_t indirectOffset)
```


TODO


**Arguments:**

 - TODO




<br/><!-- poor man's styling, just for the demo before we use a non default theme -->
***

## End { #wgpuRenderPassEncoderEnd }

```C
void wgpuRenderPassEncoderEnd(WGPURenderPassEncoder renderPassEncoder)
```


TODO


**Arguments:**

 - TODO




<br/><!-- poor man's styling, just for the demo before we use a non default theme -->
***

## EndOcclusionQuery { #wgpuRenderPassEncoderEndOcclusionQuery }

```C
void wgpuRenderPassEncoderEndOcclusionQuery(WGPURenderPassEncoder renderPassEncoder)
```


TODO


**Arguments:**

 - TODO




<br/><!-- poor man's styling, just for the demo before we use a non default theme -->
***

## ExecuteBundles { #wgpuRenderPassEncoderExecuteBundles }

```C
void wgpuRenderPassEncoderExecuteBundles(WGPURenderPassEncoder renderPassEncoder, size_t bundleCount, WGPURenderBundle const * bundles)
```


TODO


**Arguments:**

 - TODO




<br/><!-- poor man's styling, just for the demo before we use a non default theme -->
***

## InsertDebugMarker { #wgpuRenderPassEncoderInsertDebugMarker }

```C
void wgpuRenderPassEncoderInsertDebugMarker(WGPURenderPassEncoder renderPassEncoder, char const * markerLabel)
```


TODO


**Arguments:**

 - TODO




<br/><!-- poor man's styling, just for the demo before we use a non default theme -->
***

## PopDebugGroup { #wgpuRenderPassEncoderPopDebugGroup }

```C
void wgpuRenderPassEncoderPopDebugGroup(WGPURenderPassEncoder renderPassEncoder)
```


TODO


**Arguments:**

 - TODO




<br/><!-- poor man's styling, just for the demo before we use a non default theme -->
***

## PushDebugGroup { #wgpuRenderPassEncoderPushDebugGroup }

```C
void wgpuRenderPassEncoderPushDebugGroup(WGPURenderPassEncoder renderPassEncoder, char const * groupLabel)
```


TODO


**Arguments:**

 - TODO




<br/><!-- poor man's styling, just for the demo before we use a non default theme -->
***

## SetBindGroup { #wgpuRenderPassEncoderSetBindGroup }

```C
void wgpuRenderPassEncoderSetBindGroup(WGPURenderPassEncoder renderPassEncoder, uint32_t groupIndex, WGPU_NULLABLE WGPUBindGroup group, size_t dynamicOffsetCount, uint32_t const * dynamicOffsets)
```


TODO


**Arguments:**

 - TODO




<br/><!-- poor man's styling, just for the demo before we use a non default theme -->
***

## SetBlendConstant { #wgpuRenderPassEncoderSetBlendConstant }

```C
void wgpuRenderPassEncoderSetBlendConstant(WGPURenderPassEncoder renderPassEncoder, WGPUColor const * color)
```


TODO


**Arguments:**

 - TODO




<br/><!-- poor man's styling, just for the demo before we use a non default theme -->
***

## SetIndexBuffer { #wgpuRenderPassEncoderSetIndexBuffer }

```C
void wgpuRenderPassEncoderSetIndexBuffer(WGPURenderPassEncoder renderPassEncoder, WGPUBuffer buffer, WGPUIndexFormat format, uint64_t offset, uint64_t size)
```


TODO


**Arguments:**

 - TODO




<br/><!-- poor man's styling, just for the demo before we use a non default theme -->
***

## SetLabel { #wgpuRenderPassEncoderSetLabel }

```C
void wgpuRenderPassEncoderSetLabel(WGPURenderPassEncoder renderPassEncoder, char const * label)
```


TODO


**Arguments:**

 - TODO




<br/><!-- poor man's styling, just for the demo before we use a non default theme -->
***

## SetPipeline { #wgpuRenderPassEncoderSetPipeline }

```C
void wgpuRenderPassEncoderSetPipeline(WGPURenderPassEncoder renderPassEncoder, WGPURenderPipeline pipeline)
```


TODO


**Arguments:**

 - TODO




<br/><!-- poor man's styling, just for the demo before we use a non default theme -->
***

## SetScissorRect { #wgpuRenderPassEncoderSetScissorRect }

```C
void wgpuRenderPassEncoderSetScissorRect(WGPURenderPassEncoder renderPassEncoder, uint32_t x, uint32_t y, uint32_t width, uint32_t height)
```


TODO


**Arguments:**

 - TODO




<br/><!-- poor man's styling, just for the demo before we use a non default theme -->
***

## SetStencilReference { #wgpuRenderPassEncoderSetStencilReference }

```C
void wgpuRenderPassEncoderSetStencilReference(WGPURenderPassEncoder renderPassEncoder, uint32_t reference)
```


TODO


**Arguments:**

 - TODO




<br/><!-- poor man's styling, just for the demo before we use a non default theme -->
***

## SetVertexBuffer { #wgpuRenderPassEncoderSetVertexBuffer }

```C
void wgpuRenderPassEncoderSetVertexBuffer(WGPURenderPassEncoder renderPassEncoder, uint32_t slot, WGPU_NULLABLE WGPUBuffer buffer, uint64_t offset, uint64_t size)
```


TODO


**Arguments:**

 - TODO




<br/><!-- poor man's styling, just for the demo before we use a non default theme -->
***

## SetViewport { #wgpuRenderPassEncoderSetViewport }

```C
void wgpuRenderPassEncoderSetViewport(WGPURenderPassEncoder renderPassEncoder, float x, float y, float width, float height, float minDepth, float maxDepth)
```


TODO


**Arguments:**

 - TODO



