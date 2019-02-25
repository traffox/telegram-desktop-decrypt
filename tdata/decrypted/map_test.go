package decrypted

import (
	"encoding/hex"
	"fmt"
	"log"
	"sort"
)

const hexDMap0 = "d009000000000003000000582e6ab17475aa9e27000000010000689b0000000029db2dca000017cccb142f5441fb2a0d000000010003b93900000000300c72970001415c45edca91d364c281000000010003b93a00000000300c72970000032c1852f6449e9ca19b000000010003b93c00000000300c729700001b7c3894a0cb1e0917cf00000002000039b6000000000f394cd20000200c45ad764df3c6b04a00000002000060f1000000000e0803700000200cfda25d9e4542cc85000000020000f6ba000000000eab53120000041ca77ed6b767844c030000000200016030000000000d481697000017cc38eef2c8362614b8000000020001cef7000000000eb7b96a00001a3c6af556ce4d0c1ecf000000020001d8be000000000e46c2060000188c4e9317f4da2294de00000002000220c0000000000e9aa5410000143c34bc7897b8293bb40000000200035fb3000000000eabb43c0000039cbd91b5ac8a8e42940000000200035fb5000000000eabb43c00007c1c97ad5afa473db1b1000000020003bac5000000000e9c1be2000022fc608ebdcdece58a2e000000020003d008000000000eab54570000c70c69647ce2d2eaaceb000000020003d009000000000eab54570000033cf68244923026818a0000000200053465000000000eabb8520000034c79ccf70b3873732e0000000200053467000000000eabb8520000536cfbbff53429c46d6e0000000200065b77000000000eb7bd3b0000ad4cb3e34036aaacaef30000000200065b78000000000eb7bd3b0000043c87e9828da4c063b70000000400000fb900000000308b139500000abc85106b849cb944c40000000400001ba50000000030729ed300000b1c435a5e5557f1f66e0000000400003bee00000000196154b9000009bccd63270ae9e5de87000000040000aea8000000001b2ce8af0000070c15a701c008387481000000040000b1e0000000001b2cd20a0000053c3e9954d7ee5c2ad7000000040000b40f000000001b2ceb4b0000036cc047774b34f4e9e9000000040000b7e9000000001b2d047d000002ecc673334ac5f53840000000040000d72e000000001a9eec3c000006ac33e91b97e41e9e00000000040000d731000000001a9eec3c0001d8bc68215222607b45f60000000400014383000000001ac35f30000002dcf51cd8c90b9374b90000000400016b6f000000001aa0a827000003bc20f2d15a20e9b1a90000000400021ab6000000001a9a6d01000004dc96d4305313dedd870000000400024aed000000001b221a010000082c0f4e1992973d1a0d0000000400024e84000000001b2215440000043c13a6c1d10382a0c90000000400024f0e000000001b2201990000071cda2a074f86ae86340000000400024f26000000001b21f206000005fc0b711c90647f7a460000000400024fcb000000001b21e48c0000035c6eed84e853c62dc60000000400024fff000000001b221a12000006aceb27f8278b8b7f7b00000004000250ba000000001b2976910000039c2856d606d03ea0640000000400025268000000001b2219a20000068c6ea1d5edf4e6a3da0000000400025269000000001b2219a2000070dc24e50f9397646fcf0000000400025329000000001b220f830000076ca8fca44c3b2b8ded000000040002532a000000001b220f830000685cf23bab36d0431a2d000000040002546a000000001b220d4f0000052c6c0aab41736091210000000400025bcd000000001b2976b7000002cc27b2a8b6b65d63c50000000400025bce000000001b2976b70000233c555256ce1a232167000000040002ba33000000001b1ee066000005fc22ad74c2ca099908000000040002e8e1000000001b1f3b620000076c59a8ef898901fdfa000000040002ea69000000001b1f26680000076ce020bcea5f4152d8000000040002eba5000000001b1f339d000004acb3c92698566c3d0d000000040002ed76000000001b1f324a000006acd54b163d1b9bfba3000000040002edcc000000001b1f40580000046cfa853657e93625b4000000040002edcd000000001b1f405800001a5c3d70e34bab025e56000000040002ef35000000001b1f3ef90000074c25eeb34f7478ab7a000000040002ef43000000001b1f1b690000059cb91a865da8c20acc000000040002f067000000001b20568b0000068cb8d06932e289ec9c000000040002f068000000001b20568b000060ac208e1b52a3b4111e000000040002f2d3000000001b1f2a8e0000066c64892bf785330b9d000000040002f355000000001b1edb200000073c977a742020df3af7000000040002f39b000000001b204e080000066cc57fa7c3c1e7c91c000000040002f39c000000001b204e0800005a9c2c9eb56129980fe8000000040002f3d0000000001b1f2569000003fc6729c31e28748e08000000040002f3d1000000001b1f25690000310c88c22032573ed0ad000000040002f435000000001b1f06b80000053cee064a0214fdc79b000000040002f436000000001b1f06b80000352caef3b989e667fc40000000040002f456000000001b1ee146000005dcef6c7a1fd46c6d09000000040002f4d8000000001b1ef22f0000058cb10616bdef94c0f9000000040002f4f4000000001b1f0d3c0000072c69be493874cf7741000000040002f51a000000001b1f2ae90000078c8ecb51039f6383ab000000040002f51b000000001b1f2ae90000577cc9e9d1117da4e7c5000000040002f679000000001b1ef4510000070cf382b82b8fba9c48000000040002f82b000000001b1edc1c0000066c23ed803956ecff27000000040002f82c000000001b1edc1c00004ffce62c18330916bba3000000040002f91c000000001b1f2dce0000080c92f7d307b8cb1df1000000040002fa7e000000001b1ef05a000002ac2134c072cadbaa02000000040002fa7f000000001b1ef05a00001d2c34e1a27e5ee6cc42000000040002faea000000001a9f1a3c0000042c442f3797b9675fb8000000040002fb26000000001b1ee8f70000055c0583b4b8ffa1f34c000000040002fb27000000001b1ee8f700004fcc36689099f7827178000000040003069e00000000195e59220000039cca5139c9702225c700000004000413f5000000001b22212d000004bc1bdb5fc57eb4a69f00000004000418d9000000001ac363800000044cb8280e7dd0ecfcef0000000400041beb000000001ac378030000075cf671a8d6d441f2960000000400041bec000000001ac37803000068dcf3bc73631f9a275a0000000400041c60000000001b221e120000036cee4e728ea5c0045100000004000426b1000000001ac36f8a0000056cbbdd5598ff2b2cef00000004000427dd000000001ac363e4000070dc32961197d298205800000004000427de000000001ac363e40000076c0000000d735f4ef83ba7cb3b00000008b145f77302906f5200000009d4cf92b6551271c0"

func ExampleToDMap() {
	map0, _ := hex.DecodeString(hexDMap0)
	dmap, err := ReadDMap(map0)
	if err != nil {
		log.Fatal(err)
	}
	kk := []string{}
	for k := range dmap.Files {
		kk = append(kk, k)
	}
	sort.Strings(kk)
	for _, k := range kk {
		fmt.Println(k, dmap.Files[k])
	}
	// Output:
	// 00E9E14E79B19E33 3
	// 04835F5CA433376C 3
	// 04CF766E989B3FEA 3
	// 0C1721556B29FC4D 9
	// 1219063714BAA0C6 3
	// 1477FC478394EB96 3
	// 15400C5AE827E4EE 3
	// 182C463D19ACDE54 3
	// 184783800C107A51 3
	// 1B1BD374AFA5DA79 3
	// 1FD1BC8B703D7F29 3
	// 20AABDAC270C4312 3
	// 24CC6EE5E72A1E43 3
	// 25F60920377F541B 8
	// 30C448767B6DE77A 3
	// 3ABB61903381C26E 3
	// 3ABFB9B1D361B45D 3
	// 3FEACAAA63043E3B 3
	// 4368EA68F470A2AD 3
	// 460AE30D606D6582 3
	// 4924E8A8CA5B19DB 3
	// 4B52639E756358AF 3
	// 4BB3928B7987CB43 3
	// 4C449BC948B60158 3
	// 58CC2454E9D52ADF 3
	// 5C36D56B6B8A2B72 3
	// 5C7E4AD7111D9E9C 3
	// 64A7F74609C117B0 3
	// 65E520BAB43E07D3 3
	// 692F144D6D8A176F 3
	// 6CD26C358E48DEE6 3
	// 6F54B70622251286 3
	// 72E9AA57471BA6E2 3
	// 72FFCE659308DE32 3
	// 761232A1EC652555 3
	// 78DDED3135034D69 3
	// 78ED5E9EA07236DC 3
	// 7B360C4AD8289E78 3
	// 7C5222079C9315AC 3
	// 7DA2C5EE7D4599E3 3
	// 7FA3FD020247A779 3
	// 809990AC2C47DA22 3
	// 80E84782E13C9276 3
	// 84C9ABF8B28B283F 3
	// 8502892D79116923 3
	// 8717287F99098663 3
	// 8B4162638C2FEE83 3
	// 8BF5769B7973F244 3
	// 8D2514F5AECB020E 3
	// 8EF08992165BE9C2 3
	// 90D6C64DF1A7C6FE 3
	// 9A1B9E02A51D2F02 3
	// 9B4739B09C8DC15F 3
	// 9C0A28301D1C6A31 3
	// 9E9E4F43B477740C 3
	// 9F0C49FEDB61601B 3
	// A40B6C3FD467DA54 3
	// A572A9F13637CB3F 3
	// A7BA8747F43BEE52 3
	// A81862032944286F 3
	// AD3A6E4FDE5D1AE6 3
	// AFDF109898FE8A95 3
	// B3BC7AB38FE4F537 13
	// B7F7B8B8728F72BE 3
	// B91AC9E9446F2581 3
	// B97CDF4120A460EE 3
	// BA3836F93015BCE8 3
	// BECAAE2D2EC74696 3
	// C19C7E1C3C7AF75C 3
	// C43F1AFF8B4B3850 3
	// C9CE982E23960D8B 3
	// CCA02C8AD568A19B 3
	// D0A1D3792991E4F0 3
	// D0A2BF1445F241BC 3
	// D0D3C66589629C3B 3
	// D2A1340D63BAB32F 3
	// D9B033587FB29846 3
	// DA0DE37523022C88 3
	// DED8B2B3C44ACF8A 3
	// E1114B3A25B1E802 3
	// E2373783B07FCC97 3
	// E2A85ECEDCDBE806 3
	// E66F1F7555E5A534 3
	// E6D64C92435FFBBF 3
	// ED4922AD4F7139E4 3
	// F96A4BE75CF5BDB1 3
	// FC7190E1BC0A4983 3
	// FCE1C0D4EC655FA6 3
	// FCF6467939F05E42 3
	// FEC2B2FF8955DDBB 3
	// FECFCE0DD7E0828B 3
}
