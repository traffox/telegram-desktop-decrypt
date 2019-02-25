package encrypted

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"log"

	"github.com/atilaromero/telegram-desktop-decrypt/tdata"
)

const hexSettings0 string = "54444624ff4d0f000000002085b5adf5240b7d745b734fed4bf73f44d2d65231c7f334f91b862ca8228c433700000350b66feb69f32d1ac107fa2871de4500ed63a3725c19d629dd9c0d567aa2801debd0281ffa9ed8c4ea227c5cb54967485d28c7ee533e5926cde9b6643c3a43897c2456c6e1628fab8e65991a1fcdaf751bc83d4ee98bda27107b38d875eca923e188ed2b67621e83d58f68a1b4e871de6362da1b7d6081d06515bc77984e50b1816ec0317045dc8cc2e376cfeda2a80d368f5971f3206998a9b3a54e9dba5838b8dc471f5c5f95a064749480e8b983cd7a287c0042bc2ed6206a7d746d240abe54beb1f85dd377bafd8a2c161d3874f2941a0543d7bf39bdac7c22d0383bc8fe548dbcf57e8e03daeac10469ca0465b4ff0ea99cd08507ca71773ac55f5dcf0158b38e9b576f0d48ea68a47f6642b4889546fe4dae7bcba771008755094246ecee5de5af1e7d974ff33deba6f15be5fc4a7b2c5b6bf113b8a35f7c5a4444ce9a33ff8d514da7642915b27e74b4287512ed0e69dd6bac5b3bbbd1d8775b36d2fa711b29c53f9f18842b2b1cb262b6e1b20edca280eca5d8eae45ec2976ad90df84da1f504c25c9df40164e9662a38cb217570e9450e3a6c1007ff576000d5f9ab80a6abc8d05c5729ea927632cbde91a7d441fc19441d5bea0448f2bc8440799984289105eb8cde317eee70e20a629c7b78a69e95bb90e209ee5108230c84f08c7d9bf6934b31e5d2678ec671a91778da978736c009f39b6c0bc12a09b8085fd9f521d3ab6f46901cd8e544b68a0bc101279f26097447e0d28bdfc0e45479500d9557e19a82b8cdd6097d2e9c8e34f17854af20670cb5798d4550eff1c37ac5331d8ed0622c9b47c45703ba18f71ee484bfffed072e827d737f6dfc25bf1c26f2bc51d7bae7c49a30fee02414b49795b4cb320739f4d5a6601b5bba17c77474bdafa368a3d898d7df2ed413dc225ae95b1328f376ca083271a902319f729645aebf6675d3d79cb271a3f4e899a207222d51f6b1f20cacbc638baee06dbbd1a782b0724e57835de4211e0b2cd30f0c4832ec85be2e7dc08a9522cad80edfaa4d3d02e39bfa9396d0b2a4bfe29d572d49f2d509cf1a579c208be55a3c988a90f8e4df9a4a097afd03b42e47aa94cd162705d75e95068bd685746edb431bf3d8bfb0f567190402cd73eef23c1168cc56d5916afeec420d3e35858d02b28e5239fdd9ec915a86748c3285bf435d1c2c4cf3c531994db24bee29cf88f5dad74699db7c98"

func ExampleEncryptedSettings_Print() {
	settings0, _ := hex.DecodeString(hexSettings0)
	rawtdf, err := tdata.ReadRawTDF(bytes.NewReader(settings0))
	if err != nil {
		fmt.Println(err)
	}
	settings, err := ReadESettings(rawtdf)
	if err != nil {
		log.Fatal(err)
	}
	settings.Print()
	// Output:
	// salt	85b5adf5240b7d745b734fed4bf73f44d2d65231c7f334f91b862ca8228c4337
	// encrypted	b66feb69f32d1ac107fa2871de4500ed63a3725c19d629dd9c0d567aa2801debd0281ffa9ed8c4ea227c5cb54967485d28c7ee533e5926cde9b6643c3a43897c2456c6e1628fab8e65991a1fcdaf751bc83d4ee98bda27107b38d875eca923e188ed2b67621e83d58f68a1b4e871de6362da1b7d6081d06515bc77984e50b1816ec0317045dc8cc2e376cfeda2a80d368f5971f3206998a9b3a54e9dba5838b8dc471f5c5f95a064749480e8b983cd7a287c0042bc2ed6206a7d746d240abe54beb1f85dd377bafd8a2c161d3874f2941a0543d7bf39bdac7c22d0383bc8fe548dbcf57e8e03daeac10469ca0465b4ff0ea99cd08507ca71773ac55f5dcf0158b38e9b576f0d48ea68a47f6642b4889546fe4dae7bcba771008755094246ecee5de5af1e7d974ff33deba6f15be5fc4a7b2c5b6bf113b8a35f7c5a4444ce9a33ff8d514da7642915b27e74b4287512ed0e69dd6bac5b3bbbd1d8775b36d2fa711b29c53f9f18842b2b1cb262b6e1b20edca280eca5d8eae45ec2976ad90df84da1f504c25c9df40164e9662a38cb217570e9450e3a6c1007ff576000d5f9ab80a6abc8d05c5729ea927632cbde91a7d441fc19441d5bea0448f2bc8440799984289105eb8cde317eee70e20a629c7b78a69e95bb90e209ee5108230c84f08c7d9bf6934b31e5d2678ec671a91778da978736c009f39b6c0bc12a09b8085fd9f521d3ab6f46901cd8e544b68a0bc101279f26097447e0d28bdfc0e45479500d9557e19a82b8cdd6097d2e9c8e34f17854af20670cb5798d4550eff1c37ac5331d8ed0622c9b47c45703ba18f71ee484bfffed072e827d737f6dfc25bf1c26f2bc51d7bae7c49a30fee02414b49795b4cb320739f4d5a6601b5bba17c77474bdafa368a3d898d7df2ed413dc225ae95b1328f376ca083271a902319f729645aebf6675d3d79cb271a3f4e899a207222d51f6b1f20cacbc638baee06dbbd1a782b0724e57835de4211e0b2cd30f0c4832ec85be2e7dc08a9522cad80edfaa4d3d02e39bfa9396d0b2a4bfe29d572d49f2d509cf1a579c208be55a3c988a90f8e4df9a4a097afd03b42e47aa94cd162705d75e95068bd685746edb431bf3d8bfb0f567190402cd73eef23c1168cc56d5916afeec420d3e35858d02b28e5239fdd9ec915a86748c3285bf435d1c2c4cf3c531
}

func ExampleEncryptedSettings_GetKey() {
	settings0, _ := hex.DecodeString(hexSettings0)
	rawtdf, err := tdata.ReadRawTDF(bytes.NewReader(settings0))
	if err != nil {
		fmt.Println(err)
	}
	settings, err := ReadESettings(rawtdf)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(hex.EncodeToString(settings.GetKey("")))
	// Output:
	// f064d0c8be23c407c96e4d642b8274be8dd6a9a310764fdcd67aa1a7d8f7041ead8a1a39e284e53ca4689b6c28e2d56ae5158b3e2ece485f8146c474b93f9a9f1f506004106c4f822a44b8568a8b01b82a050a0f3ec760d7a967896fa4356dee39d79b4502908dbc832d15a11b86bd24f7aafbe280284b5c599aeaab304940eb470e2405dfe3bd1660d9bbc3bcc48d2a8e3ce388f3e36e4cfc6862135586173f3d6f501564dac51ab213eca397761ad50a839c218d2a01c68d36fc00ea277e3ca739996d5e89b48f5c339c93b605358afeca37d04910e60367d6acca5e6aee871aa2c9fd837f0a615402e25c70dbb59916767b38adc43e05a7b721ae150ef55d
}
func ExampleEncryptedSettings_Decrypt() {
	settings0, _ := hex.DecodeString(hexSettings0)
	rawtdf, err := tdata.ReadRawTDF(bytes.NewReader(settings0))
	if err != nil {
		fmt.Println(err)
	}
	settings, err := ReadESettings(rawtdf)
	if err != nil {
		log.Fatal(err)
	}
	key := settings.GetKey("")
	decrypted, err := settings.Decrypt(key)
	fmt.Println(hex.EncodeToString(decrypted))
	// Output:
	// 3203000000000003000000c80000003200030d4000000035000000c800000043000000c80000005000000005000000060000000000000007000000000000001d0000000000000009000000000000000a000000000000000c000000010000000d0000000000000016000000000000004a0000025cffffffff0000000d0000000100000000000001bb0000000e3134392e3135342e3137352e3530000000000000000100000001000001bb00000027323030313a306232383a663233643a663030313a303030303a303030303a303030303a30303061000000000000000200000000000001bb0000000e3134392e3135342e3136372e3531000000000000000200000001000001bb00000027323030313a303637633a303465383a663030323a303030303a303030303a303030303a30303061000000000000000300000000000001bb0000000f3134392e3135342e3137352e313030000000000000000300000001000001bb00000027323030313a306232383a663233643a663030333a303030303a303030303a303030303a30303061000000000000000400000000000001bb0000000e3134392e3135342e3136372e3931000000000000000400000001000001bb00000027323030313a303637633a303465383a663030343a303030303a303030303a303030303a30303061000000000000000400000002000001bb0000000f3134392e3135342e3136342e323530000000000000000400000003000001bb00000027323030313a303637633a303465383a663030343a303030303a303030303a303030303a30303062000000000000000500000001000001bb00000027323030313a306232383a663233663a663030353a303030303a303030303a303030303a30303061000000000000000500000010000001bb0000000d39312e3130382e35362e323030000000000000000500000000000001bb0000000d39312e3130382e35362e3137350000000000000000000000190000001a00350035003500310039003900390037003700390030003700300000004f0000000400000000ffffffff00000028000000010000004ea642b24495ee2a550000000e00000002000003c20000038800000390c674053f0000000042d807ed1721bba9ad81f709ee26
}
