package validation

import (
	"encoding/json"
	"regexp"
)

func ErrorDetection(data string) []string {
	type DocType struct {
		Type string `json:"type"`
	}
	var docType DocType
	json.Unmarshal([]byte(data), &docType)
	switch docType.Type {
	case "M11":
		return ErrorsFromM11(data)
	case "FMU76":
		return ErrorsFromFMU76(data)
	}
	return []string{"Неправильный тип документа"}
}

func ErrorsFromM11(data string) []string {
	type Data struct {
		Order           string `json:"order"`
		Organisation    string `json:"organisation"`
		StructPod       string `json:"structpod"`
		DataCreate      string `json:"datacreate"`
		OpTypeCode      string `json:"optypecode"`
		SenderStructPod string `json:"senderstructpod"`
		SenderTabelNum  string `json:"sendertabelnum"`
		ReStructPod     string `json:"restructpod"`
		ReTabelNum      string `json:"retabelnum"`
		Schet           string `json:"schet"`
		AnalCode        string `json:"analcode"`
		UchetEd         string `json:"ucheted"`
		Zatreber        string `json:"zatreber"`
		Otpusher        string `json:"otpusher"`
		Razr            string `json:"razr"`
		ThroughWho      string `json:"throughwho"`
		MatName         string `json:"matname"`
		MatNum          string `json:"matnum"`
		Charact         string `json:"charact"`
		ZavodNum        string `json:"zavodnum"`
		InvenNum        string `json:"inventnum"`
		NetworkNum      string `json:"networknum"`
		EdCode          string `json:"edcode"`
		EdName          string `json:"edname"`
		KolZatreb       string `json:"kolzatreb"`
		KolOtpush       string `json:"kolotpush"`
		Price           string `json:"price"`
		SumWithoutNds   string `json:"sumwithout"`
		PorNum          string `json:"pornum"`
		Location        string `json:"location"`
		RegNum          string `json:"regnum"`
	}

	var doc Data
	json.Unmarshal([]byte(data), &doc)
	// if err != nil {
	// 	// TODO: add logger
	// }

	//checking Fields
	errors_slice := make([]string, 1)
	//1 - Order
	re := regexp.MustCompile(`^\d{8}$`)
	if !re.MatchString(doc.Order) {
		errors_slice = append(errors_slice, `Order regex error`)
	}

	//2 - Organisation

	//3 - StructPod

	//4 - DataCreate
	re = regexp.MustCompile(`^\\d{2}\\.\\d{2}\\.\\d{4}$`)
	if !re.MatchString(doc.DataCreate) {
		errors_slice = append(errors_slice, `DataCreate regex error`)
	}

	//5 - OpTypeCode
	re = regexp.MustCompile(`^\d{3}$`)
	if !re.MatchString(doc.OpTypeCode) {
		errors_slice = append(errors_slice, `OpTypeCode regex error`)
	}
	//6 - SenderStructPod

	//8 - ReStructPod (-)

	//9 - ReTabelNum (-)

	//28 - Schet (10 цифр)
	re = regexp.MustCompile(`^\d{10}$`)
	if !re.MatchString(doc.Schet) {
		errors_slice = append(errors_slice, `Schet regex error`)
	}

	//25 - AnalCode (-)

	//10 - UchetEd (-)
	//14 - ForWho ???

	//14 - MatName

	//15 - MatNum

	re = regexp.MustCompile(`^\d{10}$`)
	if !re.MatchString(doc.MatNum) {
		errors_slice = append(errors_slice, `MatNum regex error`)
	}

	//16 - charact

	//17 - ZavodNum(-)

	//30 - InvenNum (-)

	//31 - NetworkNum (-)

	//21 - EdCode
	re = regexp.MustCompile(`^\d{3}$`)
	if !re.MatchString(doc.EdCode) {
		errors_slice = append(errors_slice, `EdCode regex error`)
	}

	//22 - EdName ( проверка на штуки)

	//20 - KolZatreb
	re = regexp.MustCompile(`^\\d+\\.\\d+$`)
	if !re.MatchString(doc.KolZatreb) {
		errors_slice = append(errors_slice, `KolZatreb regex error`)
	}

	//21 - KolOtpush
	re = regexp.MustCompile(`^\\d+\\.\\d+$`)
	if !re.MatchString(doc.KolOtpush) {
		errors_slice = append(errors_slice, `KolOtpush regex error`)
	} else if doc.KolOtpush > doc.KolZatreb {
		errors_slice = append(errors_slice, `Koltpush value error`)
	}

	//26 - Price
	re = regexp.MustCompile(`^\\d+,\\d+$`)
	if !re.MatchString(doc.Price) {
		errors_slice = append(errors_slice, `Price regex error`)
	}

	//27 - SumWithoutNds
	re = regexp.MustCompile(`^\\d+,\\d+$`)
	if !re.MatchString(doc.SumWithoutNds) {
		errors_slice = append(errors_slice, `SumWithoutNds regex error`)
	}
	//22 - PorNum(-)

	//32 - Location(-)

	//33 - RegNum(-)
	return errors_slice

}

func ErrorsFromFMU76(data string) []string {
	type DocFLU struct {
		Number        string `json:"number"`        //1
		Date          string `json:"date"`          //2
		Organization  string `json:"organization"`  //3
		Structpod     string `json:"structpod"`     //4
		Structpodmeta string `json:"structpodmeta"` //5
		Opcode        string `json:"opcode"`        //6
		Respface      string `json:"respface"`      //7
		Raskhodnap    string `json:"raskhodnap"`    //8
		Invnum        string `json:"invnum"`        //9
		Comission     string `json:"comission"`     //10
		Matmeta       string `json:"matmeta"`       //11
		Matnomnum     string `json:"matnomnum"`     //12
		Edcode        string `json:"edcode"`        //13
		Edname        string `json:"edname"`        //14
		Normkol       string `json:"normkol"`       //15
		Factkol       string `json:"factkol"`       //16
		Factprice     string `json:"factprice"`     //17
		Fuctsum       string `json:"fuctsum"`       //18
		Otklonenie    string `json:"otklonenie"`    //19
		Worktype      string `json:"worktype"`      //20
		Comment       string `json:"comment"`       //21
		Rukovod       string `json:"rukovod"`       //26
		Nositzatrat   string `json:"nositzatrat"`   //27
		Schet         string `json:"schet"`         //28
		Schetzakaz    string `json:"schetzakaz"`    //29
		Analcode      string `json:"analcode"`      //30
		Techschet     string `json:"techschet"`     //31
		Proizvzakaz   string `json:"proizvzakaz"`   //32
		Charact       string `json:"charact"`       //33
		Zavdetnum     string `json:"zavdetnum"`     //34
		Registnum     string `json:"registnum"`     //35
	}

	var doc DocFLU
	json.Unmarshal([]byte(data), &doc)
	// if err != nil {

	// }

	errors_slice := make([]string, 10)
	//1
	re1 := regexp.MustCompile(`^\d{8}$`)
	if !re1.MatchString(doc.Number) {
		errors_slice = append(errors_slice, `Regnum regex error`)
	}
	//2
	re2 := regexp.MustCompile(`^\\d{2}\\.\\d{2}\\.\\d{4}$`)
	if !re2.MatchString(doc.Date) {
		errors_slice = append(errors_slice, `Data regex error`)
	}
	//3
	//4
	//5
	//6
	re6 := regexp.MustCompile(`^\d{3}$`)
	if !re6.MatchString(doc.Opcode) {
		errors_slice = append(errors_slice, `Operation code regex error`)
	}
	//7
	//8
	//29
	re29 := regexp.MustCompile(`^\d{10}$`)
	if !re29.MatchString(doc.Schet) {
		errors_slice = append(errors_slice, `Schet regex error`)
	}
	//31
	re31 := regexp.MustCompile(`^32\d{8}$`)
	if !re31.MatchString(doc.Techschet) {
		errors_slice = append(errors_slice, `Techschet regex error`)
	}
	//32
	re32 := regexp.MustCompile(`[A-Z\d]{12}`)
	if !re32.MatchString(doc.Proizvzakaz) {
		errors_slice = append(errors_slice, `Proizvodstvennyi zakaz regex error`)
	}
	//27
	re27 := regexp.MustCompile(`\d{4}`)
	if !re27.MatchString(doc.Nositzatrat) {
		errors_slice = append(errors_slice, `Nositzatrat regex error`)
	}
	//19
	re19 := regexp.MustCompile(`^\d+$`)
	if !re19.MatchString(doc.Otklonenie) {
		errors_slice = append(errors_slice, `Otklonenie regex error`)
	}
	//17,18
	re17 := regexp.MustCompile(`^\d+$`)
	if !re17.MatchString(doc.Factprice) {
		errors_slice = append(errors_slice, `Rubl sum regex error`)
	}
	re18 := regexp.MustCompile(`^\d+$`)
	if !re18.MatchString(doc.Fuctsum) {
		errors_slice = append(errors_slice, `Fact sum regex error`)
	}
	re16 := regexp.MustCompile(`^\d+\.\d{3}$`)
	if !re16.MatchString(doc.Factkol) {
		errors_slice = append(errors_slice, `Fact kol regex error`)
	}
	re15 := regexp.MustCompile(`^\d+\.\d{3}$`)
	if !re15.MatchString(doc.Normkol) {
		errors_slice = append(errors_slice, `Normal kol regex error`)
	}
	re14 := regexp.MustCompile(`^[А-ЯЁ][а-яё]+\b`)
	if !re14.MatchString(doc.Edname) {
		errors_slice = append(errors_slice, `Ed izm name regex error`)
	}
	re13 := regexp.MustCompile(`\d{3}`)
	if !re13.MatchString(doc.Edcode) {
		errors_slice = append(errors_slice, `Ed izm code regex error`)
	}
	re12 := regexp.MustCompile(`^\d{10}$`)
	if !re12.MatchString(doc.Matnomnum) {
		errors_slice = append(errors_slice, `Nom number regex error`)
	}
	return errors_slice
}
