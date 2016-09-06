// CPPATL.cpp : Implementation of WinMain


#include "stdafx.h"
#include "resource.h"
#include "CPPATL_i.h"


using namespace ATL;


class CCPPATLModule : public ATL::CAtlExeModuleT< CCPPATLModule >
{
public :
	DECLARE_LIBID(LIBID_CPPATLLib)
	DECLARE_REGISTRY_APPID_RESOURCEID(IDR_CPPATL, "{44CCAA97-5EF4-44C3-A4B4-A7946F52CDE1}")
	};

CCPPATLModule _AtlModule;



//
extern "C" int WINAPI _tWinMain(HINSTANCE /*hInstance*/, HINSTANCE /*hPrevInstance*/, 
								LPTSTR /*lpCmdLine*/, int nShowCmd)
{
	return _AtlModule.WinMain(nShowCmd);
}

