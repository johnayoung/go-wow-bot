/* ----------------------------------------------------------------------------
 * This file was automatically generated by SWIG (http://www.swig.org).
 * Version 4.0.2
 *
 * This file is not intended to be easily readable and contains a number of
 * coding conventions designed to improve portability and efficiency. Do not make
 * changes to this file unless you know what you are doing--modify the SWIG
 * interface file instead.
 * ----------------------------------------------------------------------------- */

// source: navigation.i

package navigation

/*
#define intgo swig_intgo
typedef void *swig_voidp;

#include <stdint.h>


typedef long long intgo;
typedef unsigned long long uintgo;



typedef struct { char *p; intgo n; } _gostring_;
typedef struct { void* array; intgo len; intgo cap; } _goslice_;


extern void _wrap_Swig_free_navigation_cff4befa467ca9b5(uintptr_t arg1);
extern uintptr_t _wrap_Swig_malloc_navigation_cff4befa467ca9b5(swig_intgo arg1);
extern void _wrap_XYZ_X_set_navigation_cff4befa467ca9b5(uintptr_t arg1, float arg2);
extern float _wrap_XYZ_X_get_navigation_cff4befa467ca9b5(uintptr_t arg1);
extern void _wrap_XYZ_Y_set_navigation_cff4befa467ca9b5(uintptr_t arg1, float arg2);
extern float _wrap_XYZ_Y_get_navigation_cff4befa467ca9b5(uintptr_t arg1);
extern void _wrap_XYZ_Z_set_navigation_cff4befa467ca9b5(uintptr_t arg1, float arg2);
extern float _wrap_XYZ_Z_get_navigation_cff4befa467ca9b5(uintptr_t arg1);
extern uintptr_t _wrap_new_XYZ__SWIG_0_navigation_cff4befa467ca9b5(void);
extern uintptr_t _wrap_new_XYZ__SWIG_1_navigation_cff4befa467ca9b5(double arg1, double arg2, double arg3);
extern void _wrap_delete_XYZ_navigation_cff4befa467ca9b5(uintptr_t arg1);
extern uintptr_t _wrap_Navigation_GetInstance_navigation_cff4befa467ca9b5(void);
extern void _wrap_Navigation_Initialize_navigation_cff4befa467ca9b5(uintptr_t arg1);
extern void _wrap_Navigation_Release_navigation_cff4befa467ca9b5(uintptr_t arg1);
extern uintptr_t _wrap_Navigation_CalculatePath_navigation_cff4befa467ca9b5(uintptr_t arg1, swig_intgo arg2, uintptr_t arg3, uintptr_t arg4, _Bool arg5, swig_voidp arg6);
extern void _wrap_Navigation_FreePathArr_navigation_cff4befa467ca9b5(uintptr_t arg1, uintptr_t arg2);
extern uintptr_t _wrap_Navigation_GetMmapsPath_navigation_cff4befa467ca9b5(uintptr_t arg1);
extern uintptr_t _wrap_new_Navigation_navigation_cff4befa467ca9b5(void);
extern void _wrap_delete_Navigation_navigation_cff4befa467ca9b5(uintptr_t arg1);
extern uintptr_t _wrap_dtCustomAlloc_navigation_cff4befa467ca9b5(swig_intgo arg1, uintptr_t arg2);
extern void _wrap_dtCustomFree_navigation_cff4befa467ca9b5(uintptr_t arg1);
extern uintptr_t _wrap_new_MMapData_navigation_cff4befa467ca9b5(uintptr_t arg1);
extern void _wrap_delete_MMapData_navigation_cff4befa467ca9b5(uintptr_t arg1);
extern void _wrap_MMapData_navMesh_set_navigation_cff4befa467ca9b5(uintptr_t arg1, uintptr_t arg2);
extern uintptr_t _wrap_MMapData_navMesh_get_navigation_cff4befa467ca9b5(uintptr_t arg1);
extern void _wrap_MMapData_navMeshQueries_set_navigation_cff4befa467ca9b5(uintptr_t arg1, uintptr_t arg2);
extern uintptr_t _wrap_MMapData_navMeshQueries_get_navigation_cff4befa467ca9b5(uintptr_t arg1);
extern void _wrap_MMapData_mmapLoadedTiles_set_navigation_cff4befa467ca9b5(uintptr_t arg1, uintptr_t arg2);
extern uintptr_t _wrap_MMapData_mmapLoadedTiles_get_navigation_cff4befa467ca9b5(uintptr_t arg1);
extern void _wrap_delete_MMapManager_navigation_cff4befa467ca9b5(uintptr_t arg1);
extern void _wrap_MMapManager_hasLoadedWesternContinent_set_navigation_cff4befa467ca9b5(uintptr_t arg1, _Bool arg2);
extern _Bool _wrap_MMapManager_hasLoadedWesternContinent_get_navigation_cff4befa467ca9b5(uintptr_t arg1);
extern void _wrap_MMapManager_hasLoadedEasternContinent_set_navigation_cff4befa467ca9b5(uintptr_t arg1, _Bool arg2);
extern _Bool _wrap_MMapManager_hasLoadedEasternContinent_get_navigation_cff4befa467ca9b5(uintptr_t arg1);
extern _Bool _wrap_MMapManager_loadMap_navigation_cff4befa467ca9b5(uintptr_t arg1, swig_intgo arg2, swig_intgo arg3, swig_intgo arg4);
extern uintptr_t _wrap_MMapManager_GetNavMeshQuery_navigation_cff4befa467ca9b5(uintptr_t arg1, swig_intgo arg2, swig_intgo arg3);
extern uintptr_t _wrap_MMapManager_GetNavMesh_navigation_cff4befa467ca9b5(uintptr_t arg1, swig_intgo arg2);
extern swig_intgo _wrap_MMapManager_getLoadedMapsCount_navigation_cff4befa467ca9b5(uintptr_t arg1);
extern uintptr_t _wrap_new_MMapManager_navigation_cff4befa467ca9b5(void);
extern uintptr_t _wrap_MMapFactory_createOrGetMMapManager_navigation_cff4befa467ca9b5(void);
extern uintptr_t _wrap_new_MMapFactory_navigation_cff4befa467ca9b5(void);
extern void _wrap_delete_MMapFactory_navigation_cff4befa467ca9b5(uintptr_t arg1);
#undef intgo
*/
import "C"

import "unsafe"
import _ "runtime/cgo"
import "sync"


type _ unsafe.Pointer



var Swig_escape_always_false bool
var Swig_escape_val interface{}


type _swig_fnptr *byte
type _swig_memberptr *byte


type _ sync.Mutex

func Swig_free(arg1 uintptr) {
	_swig_i_0 := arg1
	C._wrap_Swig_free_navigation_cff4befa467ca9b5(C.uintptr_t(_swig_i_0))
}

func Swig_malloc(arg1 int) (_swig_ret uintptr) {
	var swig_r uintptr
	_swig_i_0 := arg1
	swig_r = (uintptr)(C._wrap_Swig_malloc_navigation_cff4befa467ca9b5(C.swig_intgo(_swig_i_0)))
	return swig_r
}

type SwigcptrXYZ uintptr

func (p SwigcptrXYZ) Swigcptr() uintptr {
	return (uintptr)(p)
}

func (p SwigcptrXYZ) SwigIsXYZ() {
}

func (arg1 SwigcptrXYZ) SetX(arg2 float32) {
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	C._wrap_XYZ_X_set_navigation_cff4befa467ca9b5(C.uintptr_t(_swig_i_0), C.float(_swig_i_1))
}

func (arg1 SwigcptrXYZ) GetX() (_swig_ret float32) {
	var swig_r float32
	_swig_i_0 := arg1
	swig_r = (float32)(C._wrap_XYZ_X_get_navigation_cff4befa467ca9b5(C.uintptr_t(_swig_i_0)))
	return swig_r
}

func (arg1 SwigcptrXYZ) SetY(arg2 float32) {
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	C._wrap_XYZ_Y_set_navigation_cff4befa467ca9b5(C.uintptr_t(_swig_i_0), C.float(_swig_i_1))
}

func (arg1 SwigcptrXYZ) GetY() (_swig_ret float32) {
	var swig_r float32
	_swig_i_0 := arg1
	swig_r = (float32)(C._wrap_XYZ_Y_get_navigation_cff4befa467ca9b5(C.uintptr_t(_swig_i_0)))
	return swig_r
}

func (arg1 SwigcptrXYZ) SetZ(arg2 float32) {
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	C._wrap_XYZ_Z_set_navigation_cff4befa467ca9b5(C.uintptr_t(_swig_i_0), C.float(_swig_i_1))
}

func (arg1 SwigcptrXYZ) GetZ() (_swig_ret float32) {
	var swig_r float32
	_swig_i_0 := arg1
	swig_r = (float32)(C._wrap_XYZ_Z_get_navigation_cff4befa467ca9b5(C.uintptr_t(_swig_i_0)))
	return swig_r
}

func NewXYZ__SWIG_0() (_swig_ret XYZ) {
	var swig_r XYZ
	swig_r = (XYZ)(SwigcptrXYZ(C._wrap_new_XYZ__SWIG_0_navigation_cff4befa467ca9b5()))
	return swig_r
}

func NewXYZ__SWIG_1(arg1 float64, arg2 float64, arg3 float64) (_swig_ret XYZ) {
	var swig_r XYZ
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	_swig_i_2 := arg3
	swig_r = (XYZ)(SwigcptrXYZ(C._wrap_new_XYZ__SWIG_1_navigation_cff4befa467ca9b5(C.double(_swig_i_0), C.double(_swig_i_1), C.double(_swig_i_2))))
	return swig_r
}

func NewXYZ(a ...interface{}) XYZ {
	argc := len(a)
	if argc == 0 {
		return NewXYZ__SWIG_0()
	}
	if argc == 3 {
		return NewXYZ__SWIG_1(a[0].(float64), a[1].(float64), a[2].(float64))
	}
	panic("No match for overloaded function call")
}

func DeleteXYZ(arg1 XYZ) {
	_swig_i_0 := arg1.Swigcptr()
	C._wrap_delete_XYZ_navigation_cff4befa467ca9b5(C.uintptr_t(_swig_i_0))
}

type XYZ interface {
	Swigcptr() uintptr
	SwigIsXYZ()
	SetX(arg2 float32)
	GetX() (_swig_ret float32)
	SetY(arg2 float32)
	GetY() (_swig_ret float32)
	SetZ(arg2 float32)
	GetZ() (_swig_ret float32)
}

type SwigcptrNavigation uintptr

func (p SwigcptrNavigation) Swigcptr() uintptr {
	return (uintptr)(p)
}

func (p SwigcptrNavigation) SwigIsNavigation() {
}

func NavigationGetInstance() (_swig_ret Navigation) {
	var swig_r Navigation
	swig_r = (Navigation)(SwigcptrNavigation(C._wrap_Navigation_GetInstance_navigation_cff4befa467ca9b5()))
	return swig_r
}

func (arg1 SwigcptrNavigation) Initialize() {
	_swig_i_0 := arg1
	C._wrap_Navigation_Initialize_navigation_cff4befa467ca9b5(C.uintptr_t(_swig_i_0))
}

func (arg1 SwigcptrNavigation) Release() {
	_swig_i_0 := arg1
	C._wrap_Navigation_Release_navigation_cff4befa467ca9b5(C.uintptr_t(_swig_i_0))
}

func (arg1 SwigcptrNavigation) CalculatePath(arg2 uint, arg3 XYZ, arg4 XYZ, arg5 bool, arg6 *int) (_swig_ret XYZ) {
	var swig_r XYZ
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	_swig_i_2 := arg3.Swigcptr()
	_swig_i_3 := arg4.Swigcptr()
	_swig_i_4 := arg5
	_swig_i_5 := arg6
	swig_r = (XYZ)(SwigcptrXYZ(C._wrap_Navigation_CalculatePath_navigation_cff4befa467ca9b5(C.uintptr_t(_swig_i_0), C.swig_intgo(_swig_i_1), C.uintptr_t(_swig_i_2), C.uintptr_t(_swig_i_3), C._Bool(_swig_i_4), C.swig_voidp(_swig_i_5))))
	return swig_r
}

func (arg1 SwigcptrNavigation) FreePathArr(arg2 XYZ) {
	_swig_i_0 := arg1
	_swig_i_1 := arg2.Swigcptr()
	C._wrap_Navigation_FreePathArr_navigation_cff4befa467ca9b5(C.uintptr_t(_swig_i_0), C.uintptr_t(_swig_i_1))
}

func (arg1 SwigcptrNavigation) GetMmapsPath() (_swig_ret Std_string) {
	var swig_r Std_string
	_swig_i_0 := arg1
	swig_r = (Std_string)(SwigcptrStd_string(C._wrap_Navigation_GetMmapsPath_navigation_cff4befa467ca9b5(C.uintptr_t(_swig_i_0))))
	return swig_r
}

func NewNavigation() (_swig_ret Navigation) {
	var swig_r Navigation
	swig_r = (Navigation)(SwigcptrNavigation(C._wrap_new_Navigation_navigation_cff4befa467ca9b5()))
	return swig_r
}

func DeleteNavigation(arg1 Navigation) {
	_swig_i_0 := arg1.Swigcptr()
	C._wrap_delete_Navigation_navigation_cff4befa467ca9b5(C.uintptr_t(_swig_i_0))
}

type Navigation interface {
	Swigcptr() uintptr
	SwigIsNavigation()
	Initialize()
	Release()
	CalculatePath(arg2 uint, arg3 XYZ, arg4 XYZ, arg5 bool, arg6 *int) (_swig_ret XYZ)
	FreePathArr(arg2 XYZ)
	GetMmapsPath() (_swig_ret Std_string)
}

func DtCustomAlloc(arg1 int, arg2 DtAllocHint) (_swig_ret uintptr) {
	var swig_r uintptr
	_swig_i_0 := arg1
	_swig_i_1 := arg2.Swigcptr()
	swig_r = (uintptr)(C._wrap_dtCustomAlloc_navigation_cff4befa467ca9b5(C.swig_intgo(_swig_i_0), C.uintptr_t(_swig_i_1)))
	return swig_r
}

func DtCustomFree(arg1 uintptr) {
	_swig_i_0 := arg1
	C._wrap_dtCustomFree_navigation_cff4befa467ca9b5(C.uintptr_t(_swig_i_0))
}

type SwigcptrMMapData uintptr

func (p SwigcptrMMapData) Swigcptr() uintptr {
	return (uintptr)(p)
}

func (p SwigcptrMMapData) SwigIsMMapData() {
}

func NewMMapData(arg1 DtNavMesh) (_swig_ret MMapData) {
	var swig_r MMapData
	_swig_i_0 := arg1.Swigcptr()
	swig_r = (MMapData)(SwigcptrMMapData(C._wrap_new_MMapData_navigation_cff4befa467ca9b5(C.uintptr_t(_swig_i_0))))
	return swig_r
}

func DeleteMMapData(arg1 MMapData) {
	_swig_i_0 := arg1.Swigcptr()
	C._wrap_delete_MMapData_navigation_cff4befa467ca9b5(C.uintptr_t(_swig_i_0))
}

func (arg1 SwigcptrMMapData) SetNavMesh(arg2 DtNavMesh) {
	_swig_i_0 := arg1
	_swig_i_1 := arg2.Swigcptr()
	C._wrap_MMapData_navMesh_set_navigation_cff4befa467ca9b5(C.uintptr_t(_swig_i_0), C.uintptr_t(_swig_i_1))
}

func (arg1 SwigcptrMMapData) GetNavMesh() (_swig_ret DtNavMesh) {
	var swig_r DtNavMesh
	_swig_i_0 := arg1
	swig_r = (DtNavMesh)(SwigcptrDtNavMesh(C._wrap_MMapData_navMesh_get_navigation_cff4befa467ca9b5(C.uintptr_t(_swig_i_0))))
	return swig_r
}

func (arg1 SwigcptrMMapData) SetNavMeshQueries(arg2 Std_unordered_map_Sl_unsigned_SS_int_Sc_dtNavMeshQuery_Sm__Sg_) {
	_swig_i_0 := arg1
	_swig_i_1 := arg2.Swigcptr()
	C._wrap_MMapData_navMeshQueries_set_navigation_cff4befa467ca9b5(C.uintptr_t(_swig_i_0), C.uintptr_t(_swig_i_1))
}

func (arg1 SwigcptrMMapData) GetNavMeshQueries() (_swig_ret Std_unordered_map_Sl_unsigned_SS_int_Sc_dtNavMeshQuery_Sm__Sg_) {
	var swig_r Std_unordered_map_Sl_unsigned_SS_int_Sc_dtNavMeshQuery_Sm__Sg_
	_swig_i_0 := arg1
	swig_r = (Std_unordered_map_Sl_unsigned_SS_int_Sc_dtNavMeshQuery_Sm__Sg_)(SwigcptrStd_unordered_map_Sl_unsigned_SS_int_Sc_dtNavMeshQuery_Sm__Sg_(C._wrap_MMapData_navMeshQueries_get_navigation_cff4befa467ca9b5(C.uintptr_t(_swig_i_0))))
	return swig_r
}

func (arg1 SwigcptrMMapData) SetMmapLoadedTiles(arg2 Std_unordered_map_Sl_unsigned_SS_int_Sc_dtTileRef_Sg_) {
	_swig_i_0 := arg1
	_swig_i_1 := arg2.Swigcptr()
	C._wrap_MMapData_mmapLoadedTiles_set_navigation_cff4befa467ca9b5(C.uintptr_t(_swig_i_0), C.uintptr_t(_swig_i_1))
}

func (arg1 SwigcptrMMapData) GetMmapLoadedTiles() (_swig_ret Std_unordered_map_Sl_unsigned_SS_int_Sc_dtTileRef_Sg_) {
	var swig_r Std_unordered_map_Sl_unsigned_SS_int_Sc_dtTileRef_Sg_
	_swig_i_0 := arg1
	swig_r = (Std_unordered_map_Sl_unsigned_SS_int_Sc_dtTileRef_Sg_)(SwigcptrStd_unordered_map_Sl_unsigned_SS_int_Sc_dtTileRef_Sg_(C._wrap_MMapData_mmapLoadedTiles_get_navigation_cff4befa467ca9b5(C.uintptr_t(_swig_i_0))))
	return swig_r
}

type MMapData interface {
	Swigcptr() uintptr
	SwigIsMMapData()
	SetNavMesh(arg2 DtNavMesh)
	GetNavMesh() (_swig_ret DtNavMesh)
	SetNavMeshQueries(arg2 Std_unordered_map_Sl_unsigned_SS_int_Sc_dtNavMeshQuery_Sm__Sg_)
	GetNavMeshQueries() (_swig_ret Std_unordered_map_Sl_unsigned_SS_int_Sc_dtNavMeshQuery_Sm__Sg_)
	SetMmapLoadedTiles(arg2 Std_unordered_map_Sl_unsigned_SS_int_Sc_dtTileRef_Sg_)
	GetMmapLoadedTiles() (_swig_ret Std_unordered_map_Sl_unsigned_SS_int_Sc_dtTileRef_Sg_)
}

type SwigcptrMMapManager uintptr

func (p SwigcptrMMapManager) Swigcptr() uintptr {
	return (uintptr)(p)
}

func (p SwigcptrMMapManager) SwigIsMMapManager() {
}

func DeleteMMapManager(arg1 MMapManager) {
	_swig_i_0 := arg1.Swigcptr()
	C._wrap_delete_MMapManager_navigation_cff4befa467ca9b5(C.uintptr_t(_swig_i_0))
}

func (arg1 SwigcptrMMapManager) SetHasLoadedWesternContinent(arg2 bool) {
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	C._wrap_MMapManager_hasLoadedWesternContinent_set_navigation_cff4befa467ca9b5(C.uintptr_t(_swig_i_0), C._Bool(_swig_i_1))
}

func (arg1 SwigcptrMMapManager) GetHasLoadedWesternContinent() (_swig_ret bool) {
	var swig_r bool
	_swig_i_0 := arg1
	swig_r = (bool)(C._wrap_MMapManager_hasLoadedWesternContinent_get_navigation_cff4befa467ca9b5(C.uintptr_t(_swig_i_0)))
	return swig_r
}

func (arg1 SwigcptrMMapManager) SetHasLoadedEasternContinent(arg2 bool) {
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	C._wrap_MMapManager_hasLoadedEasternContinent_set_navigation_cff4befa467ca9b5(C.uintptr_t(_swig_i_0), C._Bool(_swig_i_1))
}

func (arg1 SwigcptrMMapManager) GetHasLoadedEasternContinent() (_swig_ret bool) {
	var swig_r bool
	_swig_i_0 := arg1
	swig_r = (bool)(C._wrap_MMapManager_hasLoadedEasternContinent_get_navigation_cff4befa467ca9b5(C.uintptr_t(_swig_i_0)))
	return swig_r
}

func (arg1 SwigcptrMMapManager) LoadMap(arg2 uint, arg3 int, arg4 int) (_swig_ret bool) {
	var swig_r bool
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	_swig_i_2 := arg3
	_swig_i_3 := arg4
	swig_r = (bool)(C._wrap_MMapManager_loadMap_navigation_cff4befa467ca9b5(C.uintptr_t(_swig_i_0), C.swig_intgo(_swig_i_1), C.swig_intgo(_swig_i_2), C.swig_intgo(_swig_i_3)))
	return swig_r
}

func (arg1 SwigcptrMMapManager) GetNavMeshQuery(arg2 uint, arg3 uint) (_swig_ret DtNavMeshQuery) {
	var swig_r DtNavMeshQuery
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	_swig_i_2 := arg3
	swig_r = (DtNavMeshQuery)(SwigcptrDtNavMeshQuery(C._wrap_MMapManager_GetNavMeshQuery_navigation_cff4befa467ca9b5(C.uintptr_t(_swig_i_0), C.swig_intgo(_swig_i_1), C.swig_intgo(_swig_i_2))))
	return swig_r
}

func (arg1 SwigcptrMMapManager) GetNavMesh(arg2 uint) (_swig_ret DtNavMesh) {
	var swig_r DtNavMesh
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	swig_r = (DtNavMesh)(SwigcptrDtNavMesh(C._wrap_MMapManager_GetNavMesh_navigation_cff4befa467ca9b5(C.uintptr_t(_swig_i_0), C.swig_intgo(_swig_i_1))))
	return swig_r
}

func (arg1 SwigcptrMMapManager) GetLoadedMapsCount() (_swig_ret uint) {
	var swig_r uint
	_swig_i_0 := arg1
	swig_r = (uint)(C._wrap_MMapManager_getLoadedMapsCount_navigation_cff4befa467ca9b5(C.uintptr_t(_swig_i_0)))
	return swig_r
}

func NewMMapManager() (_swig_ret MMapManager) {
	var swig_r MMapManager
	swig_r = (MMapManager)(SwigcptrMMapManager(C._wrap_new_MMapManager_navigation_cff4befa467ca9b5()))
	return swig_r
}

type MMapManager interface {
	Swigcptr() uintptr
	SwigIsMMapManager()
	SetHasLoadedWesternContinent(arg2 bool)
	GetHasLoadedWesternContinent() (_swig_ret bool)
	SetHasLoadedEasternContinent(arg2 bool)
	GetHasLoadedEasternContinent() (_swig_ret bool)
	LoadMap(arg2 uint, arg3 int, arg4 int) (_swig_ret bool)
	GetNavMeshQuery(arg2 uint, arg3 uint) (_swig_ret DtNavMeshQuery)
	GetNavMesh(arg2 uint) (_swig_ret DtNavMesh)
	GetLoadedMapsCount() (_swig_ret uint)
}

type SwigcptrMMapFactory uintptr

func (p SwigcptrMMapFactory) Swigcptr() uintptr {
	return (uintptr)(p)
}

func (p SwigcptrMMapFactory) SwigIsMMapFactory() {
}

func MMapFactoryCreateOrGetMMapManager() (_swig_ret MMapManager) {
	var swig_r MMapManager
	swig_r = (MMapManager)(SwigcptrMMapManager(C._wrap_MMapFactory_createOrGetMMapManager_navigation_cff4befa467ca9b5()))
	return swig_r
}

func NewMMapFactory() (_swig_ret MMapFactory) {
	var swig_r MMapFactory
	swig_r = (MMapFactory)(SwigcptrMMapFactory(C._wrap_new_MMapFactory_navigation_cff4befa467ca9b5()))
	return swig_r
}

func DeleteMMapFactory(arg1 MMapFactory) {
	_swig_i_0 := arg1.Swigcptr()
	C._wrap_delete_MMapFactory_navigation_cff4befa467ca9b5(C.uintptr_t(_swig_i_0))
}

type MMapFactory interface {
	Swigcptr() uintptr
	SwigIsMMapFactory()
}


type SwigcptrDtNavMesh uintptr
type DtNavMesh interface {
	Swigcptr() uintptr;
}
func (p SwigcptrDtNavMesh) Swigcptr() uintptr {
	return uintptr(p)
}

type SwigcptrStd_unordered_map_Sl_unsigned_SS_int_Sc_dtTileRef_Sg_ uintptr
type Std_unordered_map_Sl_unsigned_SS_int_Sc_dtTileRef_Sg_ interface {
	Swigcptr() uintptr;
}
func (p SwigcptrStd_unordered_map_Sl_unsigned_SS_int_Sc_dtTileRef_Sg_) Swigcptr() uintptr {
	return uintptr(p)
}

type SwigcptrDtNavMeshQuery uintptr
type DtNavMeshQuery interface {
	Swigcptr() uintptr;
}
func (p SwigcptrDtNavMeshQuery) Swigcptr() uintptr {
	return uintptr(p)
}

type SwigcptrStd_string uintptr
type Std_string interface {
	Swigcptr() uintptr;
}
func (p SwigcptrStd_string) Swigcptr() uintptr {
	return uintptr(p)
}

type SwigcptrDtAllocHint uintptr
type DtAllocHint interface {
	Swigcptr() uintptr;
}
func (p SwigcptrDtAllocHint) Swigcptr() uintptr {
	return uintptr(p)
}

type SwigcptrStd_unordered_map_Sl_unsigned_SS_int_Sc_dtNavMeshQuery_Sm__Sg_ uintptr
type Std_unordered_map_Sl_unsigned_SS_int_Sc_dtNavMeshQuery_Sm__Sg_ interface {
	Swigcptr() uintptr;
}
func (p SwigcptrStd_unordered_map_Sl_unsigned_SS_int_Sc_dtNavMeshQuery_Sm__Sg_) Swigcptr() uintptr {
	return uintptr(p)
}

