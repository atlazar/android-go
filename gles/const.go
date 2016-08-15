// THE AUTOGENERATED LICENSE. ALL THE RIGHTS ARE RESERVED BY ROBOTS.

// WARNING: This file has automatically been generated on Sat, 28 May 2016 16:28:47 MSK.
// By http://git.io/cgogen. DO NOT EDIT.

package gles

/*
#cgo LDFLAGS: -lGLESv1_CM
#include <GLES/gl.h>
#include <stdlib.h>
#include "cgo_helpers.h"
*/
import "C"

const (
	// VersionEsCm10 as defined in GLES/gl.h:41
	VersionEsCm10 = 1
	// VersionEsCl10 as defined in GLES/gl.h:42
	VersionEsCl10 = 1
	// VersionEsCm11 as defined in GLES/gl.h:43
	VersionEsCm11 = 1
	// VersionEsCl11 as defined in GLES/gl.h:44
	VersionEsCl11 = 1
	// DepthBufferBit as defined in GLES/gl.h:47
	DepthBufferBit = 0x00000100
	// StencilBufferBit as defined in GLES/gl.h:48
	StencilBufferBit = 0x00000400
	// ColorBufferBit as defined in GLES/gl.h:49
	ColorBufferBit = 0x00004000
	// False as defined in GLES/gl.h:52
	False = 0
	// True as defined in GLES/gl.h:53
	True = 1
	// Points as defined in GLES/gl.h:56
	Points = 0x0000
	// Lines as defined in GLES/gl.h:57
	Lines = 0x0001
	// LineLoop as defined in GLES/gl.h:58
	LineLoop = 0x0002
	// LineStrip as defined in GLES/gl.h:59
	LineStrip = 0x0003
	// Triangles as defined in GLES/gl.h:60
	Triangles = 0x0004
	// TriangleStrip as defined in GLES/gl.h:61
	TriangleStrip = 0x0005
	// TriangleFan as defined in GLES/gl.h:62
	TriangleFan = 0x0006
	// Never as defined in GLES/gl.h:65
	Never = 0x0200
	// Less as defined in GLES/gl.h:66
	Less = 0x0201
	// Equal as defined in GLES/gl.h:67
	Equal = 0x0202
	// Lequal as defined in GLES/gl.h:68
	Lequal = 0x0203
	// Greater as defined in GLES/gl.h:69
	Greater = 0x0204
	// Notequal as defined in GLES/gl.h:70
	Notequal = 0x0205
	// Gequal as defined in GLES/gl.h:71
	Gequal = 0x0206
	// Always as defined in GLES/gl.h:72
	Always = 0x0207
	// Zero as defined in GLES/gl.h:75
	Zero = 0
	// One as defined in GLES/gl.h:76
	One = 1
	// SrcColor as defined in GLES/gl.h:77
	SrcColor = 0x0300
	// OneMinusSrcColor as defined in GLES/gl.h:78
	OneMinusSrcColor = 0x0301
	// SrcAlpha as defined in GLES/gl.h:79
	SrcAlpha = 0x0302
	// OneMinusSrcAlpha as defined in GLES/gl.h:80
	OneMinusSrcAlpha = 0x0303
	// DstAlpha as defined in GLES/gl.h:81
	DstAlpha = 0x0304
	// OneMinusDstAlpha as defined in GLES/gl.h:82
	OneMinusDstAlpha = 0x0305
	// DstColor as defined in GLES/gl.h:87
	DstColor = 0x0306
	// OneMinusDstColor as defined in GLES/gl.h:88
	OneMinusDstColor = 0x0307
	// SrcAlphaSaturate as defined in GLES/gl.h:89
	SrcAlphaSaturate = 0x0308
	// ClipPlane0 as defined in GLES/gl.h:96
	ClipPlane0 = 0x3000
	// ClipPlane1 as defined in GLES/gl.h:97
	ClipPlane1 = 0x3001
	// ClipPlane2 as defined in GLES/gl.h:98
	ClipPlane2 = 0x3002
	// ClipPlane3 as defined in GLES/gl.h:99
	ClipPlane3 = 0x3003
	// ClipPlane4 as defined in GLES/gl.h:100
	ClipPlane4 = 0x3004
	// ClipPlane5 as defined in GLES/gl.h:101
	ClipPlane5 = 0x3005
	// Front as defined in GLES/gl.h:115
	Front = 0x0404
	// Back as defined in GLES/gl.h:116
	Back = 0x0405
	// FrontAndBack as defined in GLES/gl.h:117
	FrontAndBack = 0x0408
	// Fog as defined in GLES/gl.h:130
	Fog = 0x0B60
	// Lighting as defined in GLES/gl.h:131
	Lighting = 0x0B50
	// Texture2d as defined in GLES/gl.h:132
	Texture2d = 0x0DE1
	// CullFace as defined in GLES/gl.h:133
	CullFace = 0x0B44
	// AlphaTest as defined in GLES/gl.h:134
	AlphaTest = 0x0BC0
	// Blend as defined in GLES/gl.h:135
	Blend = 0x0BE2
	// ColorLogicOp as defined in GLES/gl.h:136
	ColorLogicOp = 0x0BF2
	// Dither as defined in GLES/gl.h:137
	Dither = 0x0BD0
	// StencilTest as defined in GLES/gl.h:138
	StencilTest = 0x0B90
	// DepthTest as defined in GLES/gl.h:139
	DepthTest = 0x0B71
	// PointSmooth as defined in GLES/gl.h:148
	PointSmooth = 0x0B10
	// LineSmooth as defined in GLES/gl.h:149
	LineSmooth = 0x0B20
	// ScissorTest as defined in GLES/gl.h:150
	ScissorTest = 0x0C11
	// ColorMaterial as defined in GLES/gl.h:151
	ColorMaterial = 0x0B57
	// Normalize as defined in GLES/gl.h:152
	Normalize = 0x0BA1
	// RescaleNormal as defined in GLES/gl.h:153
	RescaleNormal = 0x803A
	// PolygonOffsetFill as defined in GLES/gl.h:154
	PolygonOffsetFill = 0x8037
	// VertexArray as defined in GLES/gl.h:155
	VertexArray = 0x8074
	// NormalArray as defined in GLES/gl.h:156
	NormalArray = 0x8075
	// ColorArray as defined in GLES/gl.h:157
	ColorArray = 0x8076
	// TextureCoordArray as defined in GLES/gl.h:158
	TextureCoordArray = 0x8078
	// Multisample as defined in GLES/gl.h:159
	Multisample = 0x809D
	// SampleAlphaToCoverage as defined in GLES/gl.h:160
	SampleAlphaToCoverage = 0x809E
	// SampleAlphaToOne as defined in GLES/gl.h:161
	SampleAlphaToOne = 0x809F
	// SampleCoverage as defined in GLES/gl.h:162
	SampleCoverage = 0x80A0
	// NoError as defined in GLES/gl.h:165
	NoError = 0
	// InvalidEnum as defined in GLES/gl.h:166
	InvalidEnum = 0x0500
	// InvalidValue as defined in GLES/gl.h:167
	InvalidValue = 0x0501
	// InvalidOperation as defined in GLES/gl.h:168
	InvalidOperation = 0x0502
	// StackOverflow as defined in GLES/gl.h:169
	StackOverflow = 0x0503
	// StackUnderflow as defined in GLES/gl.h:170
	StackUnderflow = 0x0504
	// OutOfMemory as defined in GLES/gl.h:171
	OutOfMemory = 0x0505
	// Exp as defined in GLES/gl.h:175
	Exp = 0x0800
	// Exp2 as defined in GLES/gl.h:176
	Exp2 = 0x0801
	// FogDensity as defined in GLES/gl.h:179
	FogDensity = 0x0B62
	// FogStart as defined in GLES/gl.h:180
	FogStart = 0x0B63
	// FogEnd as defined in GLES/gl.h:181
	FogEnd = 0x0B64
	// FogMode as defined in GLES/gl.h:182
	FogMode = 0x0B65
	// FogColor as defined in GLES/gl.h:183
	FogColor = 0x0B66
	// Cw as defined in GLES/gl.h:186
	Cw = 0x0900
	// Ccw as defined in GLES/gl.h:187
	Ccw = 0x0901
	// CurrentColor as defined in GLES/gl.h:190
	CurrentColor = 0x0B00
	// CurrentNormal as defined in GLES/gl.h:191
	CurrentNormal = 0x0B02
	// CurrentTextureCoords as defined in GLES/gl.h:192
	CurrentTextureCoords = 0x0B03
	// PointSize as defined in GLES/gl.h:193
	PointSize = 0x0B11
	// PointSizeMin as defined in GLES/gl.h:194
	PointSizeMin = 0x8126
	// PointSizeMax as defined in GLES/gl.h:195
	PointSizeMax = 0x8127
	// PointFadeThresholdSize as defined in GLES/gl.h:196
	PointFadeThresholdSize = 0x8128
	// PointDistanceAttenuation as defined in GLES/gl.h:197
	PointDistanceAttenuation = 0x8129
	// SmoothPointSizeRange as defined in GLES/gl.h:198
	SmoothPointSizeRange = 0x0B12
	// LineWidth as defined in GLES/gl.h:199
	LineWidth = 0x0B21
	// SmoothLineWidthRange as defined in GLES/gl.h:200
	SmoothLineWidthRange = 0x0B22
	// AliasedPointSizeRange as defined in GLES/gl.h:201
	AliasedPointSizeRange = 0x846D
	// AliasedLineWidthRange as defined in GLES/gl.h:202
	AliasedLineWidthRange = 0x846E
	// CullFaceMode as defined in GLES/gl.h:203
	CullFaceMode = 0x0B45
	// FrontFace as defined in GLES/gl.h:204
	FrontFace = 0x0B46
	// ShadeModel as defined in GLES/gl.h:205
	ShadeModel = 0x0B54
	// DepthRange as defined in GLES/gl.h:206
	DepthRange = 0x0B70
	// DepthWritemask as defined in GLES/gl.h:207
	DepthWritemask = 0x0B72
	// DepthClearValue as defined in GLES/gl.h:208
	DepthClearValue = 0x0B73
	// DepthFunc as defined in GLES/gl.h:209
	DepthFunc = 0x0B74
	// StencilClearValue as defined in GLES/gl.h:210
	StencilClearValue = 0x0B91
	// StencilFunc as defined in GLES/gl.h:211
	StencilFunc = 0x0B92
	// StencilValueMask as defined in GLES/gl.h:212
	StencilValueMask = 0x0B93
	// StencilFail as defined in GLES/gl.h:213
	StencilFail = 0x0B94
	// StencilPassDepthFail as defined in GLES/gl.h:214
	StencilPassDepthFail = 0x0B95
	// StencilPassDepthPass as defined in GLES/gl.h:215
	StencilPassDepthPass = 0x0B96
	// StencilRef as defined in GLES/gl.h:216
	StencilRef = 0x0B97
	// StencilWritemask as defined in GLES/gl.h:217
	StencilWritemask = 0x0B98
	// MatrixMode as defined in GLES/gl.h:218
	MatrixMode = 0x0BA0
	// Viewport as defined in GLES/gl.h:219
	Viewport = 0x0BA2
	// ModelviewStackDepth as defined in GLES/gl.h:220
	ModelviewStackDepth = 0x0BA3
	// ProjectionStackDepth as defined in GLES/gl.h:221
	ProjectionStackDepth = 0x0BA4
	// TextureStackDepth as defined in GLES/gl.h:222
	TextureStackDepth = 0x0BA5
	// ModelviewMatrix as defined in GLES/gl.h:223
	ModelviewMatrix = 0x0BA6
	// ProjectionMatrix as defined in GLES/gl.h:224
	ProjectionMatrix = 0x0BA7
	// TextureMatrix as defined in GLES/gl.h:225
	TextureMatrix = 0x0BA8
	// AlphaTestFunc as defined in GLES/gl.h:226
	AlphaTestFunc = 0x0BC1
	// AlphaTestRef as defined in GLES/gl.h:227
	AlphaTestRef = 0x0BC2
	// BlendDst as defined in GLES/gl.h:228
	BlendDst = 0x0BE0
	// BlendSrc as defined in GLES/gl.h:229
	BlendSrc = 0x0BE1
	// LogicOpMode as defined in GLES/gl.h:230
	LogicOpMode = 0x0BF0
	// ScissorBox as defined in GLES/gl.h:231
	ScissorBox = 0x0C10
	// ColorClearValue as defined in GLES/gl.h:233
	ColorClearValue = 0x0C22
	// ColorWritemask as defined in GLES/gl.h:234
	ColorWritemask = 0x0C23
	// UnpackAlignment as defined in GLES/gl.h:235
	UnpackAlignment = 0x0CF5
	// PackAlignment as defined in GLES/gl.h:236
	PackAlignment = 0x0D05
	// MaxLights as defined in GLES/gl.h:237
	MaxLights = 0x0D31
	// MaxClipPlanes as defined in GLES/gl.h:238
	MaxClipPlanes = 0x0D32
	// MaxTextureSize as defined in GLES/gl.h:239
	MaxTextureSize = 0x0D33
	// MaxModelviewStackDepth as defined in GLES/gl.h:240
	MaxModelviewStackDepth = 0x0D36
	// MaxProjectionStackDepth as defined in GLES/gl.h:241
	MaxProjectionStackDepth = 0x0D38
	// MaxTextureStackDepth as defined in GLES/gl.h:242
	MaxTextureStackDepth = 0x0D39
	// MaxViewportDims as defined in GLES/gl.h:243
	MaxViewportDims = 0x0D3A
	// MaxTextureUnits as defined in GLES/gl.h:244
	MaxTextureUnits = 0x84E2
	// SubpixelBits as defined in GLES/gl.h:245
	SubpixelBits = 0x0D50
	// RedBits as defined in GLES/gl.h:246
	RedBits = 0x0D52
	// GreenBits as defined in GLES/gl.h:247
	GreenBits = 0x0D53
	// BlueBits as defined in GLES/gl.h:248
	BlueBits = 0x0D54
	// AlphaBits as defined in GLES/gl.h:249
	AlphaBits = 0x0D55
	// DepthBits as defined in GLES/gl.h:250
	DepthBits = 0x0D56
	// StencilBits as defined in GLES/gl.h:251
	StencilBits = 0x0D57
	// PolygonOffsetUnits as defined in GLES/gl.h:252
	PolygonOffsetUnits = 0x2A00
	// PolygonOffsetFactor as defined in GLES/gl.h:254
	PolygonOffsetFactor = 0x8038
	// TextureBinding2d as defined in GLES/gl.h:255
	TextureBinding2d = 0x8069
	// VertexArraySize as defined in GLES/gl.h:256
	VertexArraySize = 0x807A
	// VertexArrayType as defined in GLES/gl.h:257
	VertexArrayType = 0x807B
	// VertexArrayStride as defined in GLES/gl.h:258
	VertexArrayStride = 0x807C
	// NormalArrayType as defined in GLES/gl.h:259
	NormalArrayType = 0x807E
	// NormalArrayStride as defined in GLES/gl.h:260
	NormalArrayStride = 0x807F
	// ColorArraySize as defined in GLES/gl.h:261
	ColorArraySize = 0x8081
	// ColorArrayType as defined in GLES/gl.h:262
	ColorArrayType = 0x8082
	// ColorArrayStride as defined in GLES/gl.h:263
	ColorArrayStride = 0x8083
	// TextureCoordArraySize as defined in GLES/gl.h:264
	TextureCoordArraySize = 0x8088
	// TextureCoordArrayType as defined in GLES/gl.h:265
	TextureCoordArrayType = 0x8089
	// TextureCoordArrayStride as defined in GLES/gl.h:266
	TextureCoordArrayStride = 0x808A
	// VertexArrayPointer as defined in GLES/gl.h:267
	VertexArrayPointer = 0x808E
	// NormalArrayPointer as defined in GLES/gl.h:268
	NormalArrayPointer = 0x808F
	// ColorArrayPointer as defined in GLES/gl.h:269
	ColorArrayPointer = 0x8090
	// TextureCoordArrayPointer as defined in GLES/gl.h:270
	TextureCoordArrayPointer = 0x8092
	// SampleBuffers as defined in GLES/gl.h:271
	SampleBuffers = 0x80A8
	// Samples as defined in GLES/gl.h:272
	Samples = 0x80A9
	// SampleCoverageValue as defined in GLES/gl.h:273
	SampleCoverageValue = 0x80AA
	// SampleCoverageInvert as defined in GLES/gl.h:274
	SampleCoverageInvert = 0x80AB
	// NumCompressedTextureFormats as defined in GLES/gl.h:282
	NumCompressedTextureFormats = 0x86A2
	// CompressedTextureFormats as defined in GLES/gl.h:283
	CompressedTextureFormats = 0x86A3
	// DontCare as defined in GLES/gl.h:286
	DontCare = 0x1100
	// Fastest as defined in GLES/gl.h:287
	Fastest = 0x1101
	// Nicest as defined in GLES/gl.h:288
	Nicest = 0x1102
	// PerspectiveCorrectionHint as defined in GLES/gl.h:291
	PerspectiveCorrectionHint = 0x0C50
	// PointSmoothHint as defined in GLES/gl.h:292
	PointSmoothHint = 0x0C51
	// LineSmoothHint as defined in GLES/gl.h:293
	LineSmoothHint = 0x0C52
	// FogHint as defined in GLES/gl.h:294
	FogHint = 0x0C54
	// GenerateMipmapHint as defined in GLES/gl.h:295
	GenerateMipmapHint = 0x8192
	// LightModelAmbient as defined in GLES/gl.h:298
	LightModelAmbient = 0x0B53
	// LightModelTwoSide as defined in GLES/gl.h:299
	LightModelTwoSide = 0x0B52
	// Ambient as defined in GLES/gl.h:302
	Ambient = 0x1200
	// Diffuse as defined in GLES/gl.h:303
	Diffuse = 0x1201
	// Specular as defined in GLES/gl.h:304
	Specular = 0x1202
	// Position as defined in GLES/gl.h:305
	Position = 0x1203
	// SpotDirection as defined in GLES/gl.h:306
	SpotDirection = 0x1204
	// SpotExponent as defined in GLES/gl.h:307
	SpotExponent = 0x1205
	// SpotCutoff as defined in GLES/gl.h:308
	SpotCutoff = 0x1206
	// ConstantAttenuation as defined in GLES/gl.h:309
	ConstantAttenuation = 0x1207
	// LinearAttenuation as defined in GLES/gl.h:310
	LinearAttenuation = 0x1208
	// QuadraticAttenuation as defined in GLES/gl.h:311
	QuadraticAttenuation = 0x1209
	// ByteType as defined in GLES/gl.h:314
	ByteType = 0x1400
	// UnsignedByteType as defined in GLES/gl.h:315
	UnsignedByteType = 0x1401
	// ShortType as defined in GLES/gl.h:316
	ShortType = 0x1402
	// UnsignedShortType as defined in GLES/gl.h:317
	UnsignedShortType = 0x1403
	// FloatType as defined in GLES/gl.h:318
	FloatType = 0x1406
	// FixedType as defined in GLES/gl.h:319
	FixedType = 0x140C
	// Clear as defined in GLES/gl.h:322
	Clear = 0x1500
	// And as defined in GLES/gl.h:323
	And = 0x1501
	// AndReverse as defined in GLES/gl.h:324
	AndReverse = 0x1502
	// Copy as defined in GLES/gl.h:325
	Copy = 0x1503
	// AndInverted as defined in GLES/gl.h:326
	AndInverted = 0x1504
	// Noop as defined in GLES/gl.h:327
	Noop = 0x1505
	// Xor as defined in GLES/gl.h:328
	Xor = 0x1506
	// Or as defined in GLES/gl.h:329
	Or = 0x1507
	// Nor as defined in GLES/gl.h:330
	Nor = 0x1508
	// Equiv as defined in GLES/gl.h:331
	Equiv = 0x1509
	// Invert as defined in GLES/gl.h:332
	Invert = 0x150A
	// OrReverse as defined in GLES/gl.h:333
	OrReverse = 0x150B
	// CopyInverted as defined in GLES/gl.h:334
	CopyInverted = 0x150C
	// OrInverted as defined in GLES/gl.h:335
	OrInverted = 0x150D
	// Nand as defined in GLES/gl.h:336
	Nand = 0x150E
	// Set as defined in GLES/gl.h:337
	Set = 0x150F
	// Emission as defined in GLES/gl.h:343
	Emission = 0x1600
	// Shininess as defined in GLES/gl.h:344
	Shininess = 0x1601
	// AmbientAndDiffuse as defined in GLES/gl.h:345
	AmbientAndDiffuse = 0x1602
	// Modelview as defined in GLES/gl.h:351
	Modelview = 0x1700
	// Projection as defined in GLES/gl.h:352
	Projection = 0x1701
	// Texture as defined in GLES/gl.h:353
	Texture = 0x1702
	// Alpha as defined in GLES/gl.h:362
	Alpha = 0x1906
	// Rgb as defined in GLES/gl.h:363
	Rgb = 0x1907
	// Rgba as defined in GLES/gl.h:364
	Rgba = 0x1908
	// Luminance as defined in GLES/gl.h:365
	Luminance = 0x1909
	// LuminanceAlpha as defined in GLES/gl.h:366
	LuminanceAlpha = 0x190A
	// UnsignedShort4444 as defined in GLES/gl.h:374
	UnsignedShort4444 = 0x8033
	// UnsignedShort5551 as defined in GLES/gl.h:375
	UnsignedShort5551 = 0x8034
	// UnsignedShort565 as defined in GLES/gl.h:376
	UnsignedShort565 = 0x8363
	// Flat as defined in GLES/gl.h:379
	Flat = 0x1D00
	// Smooth as defined in GLES/gl.h:380
	Smooth = 0x1D01
	// Keep as defined in GLES/gl.h:394
	Keep = 0x1E00
	// Replace as defined in GLES/gl.h:395
	Replace = 0x1E01
	// Incr as defined in GLES/gl.h:396
	Incr = 0x1E02
	// Decr as defined in GLES/gl.h:397
	Decr = 0x1E03
	// Vendor as defined in GLES/gl.h:401
	Vendor = 0x1F00
	// Renderer as defined in GLES/gl.h:402
	Renderer = 0x1F01
	// Version as defined in GLES/gl.h:403
	Version = 0x1F02
	// Extensions as defined in GLES/gl.h:404
	Extensions = 0x1F03
	// Modulate as defined in GLES/gl.h:413
	Modulate = 0x2100
	// Decal as defined in GLES/gl.h:414
	Decal = 0x2101
	// Add as defined in GLES/gl.h:416
	Add = 0x0104
	// TextureEnvMode as defined in GLES/gl.h:420
	TextureEnvMode = 0x2200
	// TextureEnvColor as defined in GLES/gl.h:421
	TextureEnvColor = 0x2201
	// TextureEnv as defined in GLES/gl.h:424
	TextureEnv = 0x2300
	// Nearest as defined in GLES/gl.h:427
	Nearest = 0x2600
	// Linear as defined in GLES/gl.h:428
	Linear = 0x2601
	// NearestMipmapNearest as defined in GLES/gl.h:433
	NearestMipmapNearest = 0x2700
	// LinearMipmapNearest as defined in GLES/gl.h:434
	LinearMipmapNearest = 0x2701
	// NearestMipmapLinear as defined in GLES/gl.h:435
	NearestMipmapLinear = 0x2702
	// LinearMipmapLinear as defined in GLES/gl.h:436
	LinearMipmapLinear = 0x2703
	// TextureMagFilter as defined in GLES/gl.h:439
	TextureMagFilter = 0x2800
	// TextureMinFilter as defined in GLES/gl.h:440
	TextureMinFilter = 0x2801
	// TextureWrapS as defined in GLES/gl.h:441
	TextureWrapS = 0x2802
	// TextureWrapT as defined in GLES/gl.h:442
	TextureWrapT = 0x2803
	// GenerateMipmap as defined in GLES/gl.h:443
	GenerateMipmap = 0x8191
	// Texture0 as defined in GLES/gl.h:449
	Texture0 = 0x84C0
	// Texture1 as defined in GLES/gl.h:450
	Texture1 = 0x84C1
	// Texture2 as defined in GLES/gl.h:451
	Texture2 = 0x84C2
	// Texture3 as defined in GLES/gl.h:452
	Texture3 = 0x84C3
	// Texture4 as defined in GLES/gl.h:453
	Texture4 = 0x84C4
	// Texture5 as defined in GLES/gl.h:454
	Texture5 = 0x84C5
	// Texture6 as defined in GLES/gl.h:455
	Texture6 = 0x84C6
	// Texture7 as defined in GLES/gl.h:456
	Texture7 = 0x84C7
	// Texture8 as defined in GLES/gl.h:457
	Texture8 = 0x84C8
	// Texture9 as defined in GLES/gl.h:458
	Texture9 = 0x84C9
	// Texture10 as defined in GLES/gl.h:459
	Texture10 = 0x84CA
	// Texture11 as defined in GLES/gl.h:460
	Texture11 = 0x84CB
	// Texture12 as defined in GLES/gl.h:461
	Texture12 = 0x84CC
	// Texture13 as defined in GLES/gl.h:462
	Texture13 = 0x84CD
	// Texture14 as defined in GLES/gl.h:463
	Texture14 = 0x84CE
	// Texture15 as defined in GLES/gl.h:464
	Texture15 = 0x84CF
	// Texture16 as defined in GLES/gl.h:465
	Texture16 = 0x84D0
	// Texture17 as defined in GLES/gl.h:466
	Texture17 = 0x84D1
	// Texture18 as defined in GLES/gl.h:467
	Texture18 = 0x84D2
	// Texture19 as defined in GLES/gl.h:468
	Texture19 = 0x84D3
	// Texture20 as defined in GLES/gl.h:469
	Texture20 = 0x84D4
	// Texture21 as defined in GLES/gl.h:470
	Texture21 = 0x84D5
	// Texture22 as defined in GLES/gl.h:471
	Texture22 = 0x84D6
	// Texture23 as defined in GLES/gl.h:472
	Texture23 = 0x84D7
	// Texture24 as defined in GLES/gl.h:473
	Texture24 = 0x84D8
	// Texture25 as defined in GLES/gl.h:474
	Texture25 = 0x84D9
	// Texture26 as defined in GLES/gl.h:475
	Texture26 = 0x84DA
	// Texture27 as defined in GLES/gl.h:476
	Texture27 = 0x84DB
	// Texture28 as defined in GLES/gl.h:477
	Texture28 = 0x84DC
	// Texture29 as defined in GLES/gl.h:478
	Texture29 = 0x84DD
	// Texture30 as defined in GLES/gl.h:479
	Texture30 = 0x84DE
	// Texture31 as defined in GLES/gl.h:480
	Texture31 = 0x84DF
	// ActiveTexture as defined in GLES/gl.h:481
	ActiveTexture = 0x84E0
	// ClientActiveTexture as defined in GLES/gl.h:482
	ClientActiveTexture = 0x84E1
	// Repeat as defined in GLES/gl.h:485
	Repeat = 0x2901
	// ClampToEdge as defined in GLES/gl.h:486
	ClampToEdge = 0x812F
	// Light0 as defined in GLES/gl.h:495
	Light0 = 0x4000
	// Light1 as defined in GLES/gl.h:496
	Light1 = 0x4001
	// Light2 as defined in GLES/gl.h:497
	Light2 = 0x4002
	// Light3 as defined in GLES/gl.h:498
	Light3 = 0x4003
	// Light4 as defined in GLES/gl.h:499
	Light4 = 0x4004
	// Light5 as defined in GLES/gl.h:500
	Light5 = 0x4005
	// Light6 as defined in GLES/gl.h:501
	Light6 = 0x4006
	// Light7 as defined in GLES/gl.h:502
	Light7 = 0x4007
	// ArrayBuffer as defined in GLES/gl.h:505
	ArrayBuffer = 0x8892
	// ElementArrayBuffer as defined in GLES/gl.h:506
	ElementArrayBuffer = 0x8893
	// ArrayBufferBinding as defined in GLES/gl.h:508
	ArrayBufferBinding = 0x8894
	// ElementArrayBufferBinding as defined in GLES/gl.h:509
	ElementArrayBufferBinding = 0x8895
	// VertexArrayBufferBinding as defined in GLES/gl.h:510
	VertexArrayBufferBinding = 0x8896
	// NormalArrayBufferBinding as defined in GLES/gl.h:511
	NormalArrayBufferBinding = 0x8897
	// ColorArrayBufferBinding as defined in GLES/gl.h:512
	ColorArrayBufferBinding = 0x8898
	// TextureCoordArrayBufferBinding as defined in GLES/gl.h:513
	TextureCoordArrayBufferBinding = 0x889A
	// StaticDraw as defined in GLES/gl.h:515
	StaticDraw = 0x88E4
	// DynamicDraw as defined in GLES/gl.h:516
	DynamicDraw = 0x88E8
	// BufferSize as defined in GLES/gl.h:518
	BufferSize = 0x8764
	// BufferUsage as defined in GLES/gl.h:519
	BufferUsage = 0x8765
	// Subtract as defined in GLES/gl.h:522
	Subtract = 0x84E7
	// Combine as defined in GLES/gl.h:523
	Combine = 0x8570
	// CombineRgb as defined in GLES/gl.h:524
	CombineRgb = 0x8571
	// CombineAlpha as defined in GLES/gl.h:525
	CombineAlpha = 0x8572
	// RgbScale as defined in GLES/gl.h:526
	RgbScale = 0x8573
	// AddSigned as defined in GLES/gl.h:527
	AddSigned = 0x8574
	// Interpolate as defined in GLES/gl.h:528
	Interpolate = 0x8575
	// Constant as defined in GLES/gl.h:529
	Constant = 0x8576
	// PrimaryColor as defined in GLES/gl.h:530
	PrimaryColor = 0x8577
	// Previous as defined in GLES/gl.h:531
	Previous = 0x8578
	// Operand0Rgb as defined in GLES/gl.h:532
	Operand0Rgb = 0x8590
	// Operand1Rgb as defined in GLES/gl.h:533
	Operand1Rgb = 0x8591
	// Operand2Rgb as defined in GLES/gl.h:534
	Operand2Rgb = 0x8592
	// Operand0Alpha as defined in GLES/gl.h:535
	Operand0Alpha = 0x8598
	// Operand1Alpha as defined in GLES/gl.h:536
	Operand1Alpha = 0x8599
	// Operand2Alpha as defined in GLES/gl.h:537
	Operand2Alpha = 0x859A
	// AlphaScale as defined in GLES/gl.h:539
	AlphaScale = 0x0D1C
	// Src0Rgb as defined in GLES/gl.h:541
	Src0Rgb = 0x8580
	// Src1Rgb as defined in GLES/gl.h:542
	Src1Rgb = 0x8581
	// Src2Rgb as defined in GLES/gl.h:543
	Src2Rgb = 0x8582
	// Src0Alpha as defined in GLES/gl.h:544
	Src0Alpha = 0x8588
	// Src1Alpha as defined in GLES/gl.h:545
	Src1Alpha = 0x8589
	// Src2Alpha as defined in GLES/gl.h:546
	Src2Alpha = 0x858A
	// Dot3Rgb as defined in GLES/gl.h:548
	Dot3Rgb = 0x86AE
	// Dot3Rgba as defined in GLES/gl.h:549
	Dot3Rgba = 0x86AF
	// ImplementationColorReadTypeOes as defined in GLES/gl.h:557
	ImplementationColorReadTypeOes = 0x8B9A
	// ImplementationColorReadFormatOes as defined in GLES/gl.h:558
	ImplementationColorReadFormatOes = 0x8B9B
	// Palette4Rgb8Oes as defined in GLES/gl.h:563
	Palette4Rgb8Oes = 0x8B90
	// Palette4Rgba8Oes as defined in GLES/gl.h:564
	Palette4Rgba8Oes = 0x8B91
	// Palette4R5G6B5Oes as defined in GLES/gl.h:565
	Palette4R5G6B5Oes = 0x8B92
	// Palette4Rgba4Oes as defined in GLES/gl.h:566
	Palette4Rgba4Oes = 0x8B93
	// Palette4Rgb5A1Oes as defined in GLES/gl.h:567
	Palette4Rgb5A1Oes = 0x8B94
	// Palette8Rgb8Oes as defined in GLES/gl.h:568
	Palette8Rgb8Oes = 0x8B95
	// Palette8Rgba8Oes as defined in GLES/gl.h:569
	Palette8Rgba8Oes = 0x8B96
	// Palette8R5G6B5Oes as defined in GLES/gl.h:570
	Palette8R5G6B5Oes = 0x8B97
	// Palette8Rgba4Oes as defined in GLES/gl.h:571
	Palette8Rgba4Oes = 0x8B98
	// Palette8Rgb5A1Oes as defined in GLES/gl.h:572
	Palette8Rgb5A1Oes = 0x8B99
	// PointSizeArrayOes as defined in GLES/gl.h:577
	PointSizeArrayOes = 0x8B9C
	// PointSizeArrayTypeOes as defined in GLES/gl.h:578
	PointSizeArrayTypeOes = 0x898A
	// PointSizeArrayStrideOes as defined in GLES/gl.h:579
	PointSizeArrayStrideOes = 0x898B
	// PointSizeArrayPointerOes as defined in GLES/gl.h:580
	PointSizeArrayPointerOes = 0x898C
	// PointSizeArrayBufferBindingOes as defined in GLES/gl.h:581
	PointSizeArrayBufferBindingOes = 0x8B9F
	// PointSpriteOes as defined in GLES/gl.h:586
	PointSpriteOes = 0x8861
	// CoordReplaceOes as defined in GLES/gl.h:587
	CoordReplaceOes = 0x8862
	// OesReadFormat as defined in GLES/gl.h:746
	OesReadFormat = 1
	// OesCompressedPalettedTexture as defined in GLES/gl.h:751
	OesCompressedPalettedTexture = 1
	// OesPointSizeArray as defined in GLES/gl.h:756
	OesPointSizeArray = 1
	// OesPointSprite as defined in GLES/gl.h:762
	OesPointSprite = 1
)