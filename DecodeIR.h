#ifndef _MY_PACKAGE_FOO_H_
#define _MY_PACKAGE_FOO_H_

#ifdef __cplusplus
extern "C" {
#endif
    void  DecodeIR
    (	unsigned int* Context,
        int* TpaiBursts, 	
        int TiFreq, 
        int TiSingleBurstCount, 
        int TiRepeatBurstCount, 
        char* TsProtocol, 
        int* TiDevice, 
        int* TiSubDevice, 
        int* TiOBC, 
        int TaiHex[4], 
        char* TsMisc, 
        char* TsError);

    //int _stdcall ProtocolSupportLevel(char * TsProtocol);

    //void _stdcall EnumerateProtocols(int iProtocolNumber, char* TsProtocol);

    //void _stdcall Version(char *Result);
    int get_hex_code(char *pronto, int* frequency, int* intro_length, int* rep_length, int data[]);
#else
    void  DecodeIR
    (	unsigned int* Context,
        int* TpaiBursts,
        int TiFreq,
        int TiSingleBurstCount,
        int TiRepeatBurstCount,
        char* TsProtocol,
        int* TiDevice,
        int* TiSubDevice,
        int* TiOBC,
        int TaiHex[4],
        char* TsMisc,
        char* TsError);
    int get_hex_code(char *pronto, int* frequency, int* intro_length, int* rep_length, int data[]);

#ifdef __cplusplus
}
#endif

#endif