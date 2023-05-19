package controller

import (
	"fmt"
	"microservice_spreadsheet/pkg/entity"
	"microservice_spreadsheet/pkg/security"
	"microservice_spreadsheet/pkg/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/xuri/excelize/v2"
)

// GetSubmissiveRemarks Função que chama método GetSubmissiveRemark do service e retorna json com lista
func GenerateSpreadSheet(c *gin.Context, service service.RemarkServiceInterface) {
	// Pega permissões do usuário
	permissions, err := security.GetPermissions(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	// Pega id passada como token na rota
	id, err := strconv.Atoi(fmt.Sprint(permissions["userID"]))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Cria variável do tipo user (inicialmente vazia)
	var filter *entity.RemarkFilter

	// Converte json em user
	err = c.ShouldBind(&filter)
	// Verifica se tem erro
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx := c.Request.Context()

	remarks, err := service.RemarkFilter(&id, filter, ctx)
	if err != nil {
		JSONMessenger(c, http.StatusInternalServerError, c.Request.URL.Path, err)
		return
	}

	CreateExcel(remarks)

	c.JSON(http.StatusOK, gin.H{
		"message": "Excel of remarks successfully generated",
	})

}

// JSONMessenger é responsável por enviar uma mensagem JSON de erro para o cliente.
func JSONMessenger(c *gin.Context, status int, path string, err error) {
	errorMessage := ""
	if err != nil {
		errorMessage = err.Error()
	}
	c.JSON(status, gin.H{
		"status":  status,
		"message": errorMessage,
		"error":   err,
		"path":    path,
	})
}

func CreateExcel(list_remark *entity.RemarkList) {
	// Cria arquivo de planilha Excel vazia e atribui a variável f
	f := excelize.NewFile()

	// Cria novo estilo para o cabeçalho da planilha (formatando planilha)
	styleHeader, err := f.NewStyle(&excelize.Style{
		// Borda completa na planilha (left, top, bottom e right)
		Border: []excelize.Border{
			{Type: "left", Color: "#000000", Style: 2},
			{Type: "top", Color: "#000000", Style: 2},
			{Type: "bottom", Color: "#000000", Style: 2},
			{Type: "right", Color: "#000000", Style: 2},
		},
		// Preenchimento da célula (deixa colorida)
		Fill: excelize.Fill{
			Type:    "pattern",
			Pattern: 1,
			Color:   []string{"#C0C0C0"},
		},
		// Alinhamneto do texto na célula (centralizado e não dar quebra de linha)
		Alignment: &excelize.Alignment{
			Horizontal: "center",
			WrapText:   true,
		},
		// Fonte em negrito e colorida
		Font: &excelize.Font{
			Bold:  true,
			Color: "#000000",
		},
	})
	// Verifica se teve algum erro ao criar esse novo estilo
	if err != nil {
		return
	}

	// Cria novo estilo para o corpo da planilha (formatando planilha)
	styleBody, err := f.NewStyle(&excelize.Style{
		// Borda completa na planilha (left, top, bottom e right)
		Border: []excelize.Border{
			{Type: "left", Color: "#000000", Style: 2},
			{Type: "top", Color: "#000000", Style: 2},
			{Type: "bottom", Color: "#000000", Style: 2},
			{Type: "right", Color: "#000000", Style: 2},
		},
		// Preenchimento da célula (deixa colorida)
		Fill: excelize.Fill{
			Type:    "pattern",
			Pattern: 1,
			Color:   []string{"#FFFF"},
		},
		// Alinhamneto do texto na célula (centralizado e não dar quebra de linha)
		Alignment: &excelize.Alignment{
			Horizontal: "center",
			WrapText:   true,
		},
	})
	// Verifica se teve algum erro ao criar esse novo estilo
	if err != nil {
		return
	}

	// Colocando os Sheet1 no cabeçalho da planilha
	f.SetCellValue("Sheet1", "A1", "ID")
	f.SetCellValue("Sheet1", "B1", "Name")
	f.SetCellValue("Sheet1", "C1", "Text")
	f.SetCellValue("Sheet1", "D1", "Date")
	f.SetCellValue("Sheet1", "E1", "Return")
	f.SetCellValue("Sheet1", "F1", "Subject Title")
	f.SetCellValue("Sheet1", "G1", "Client")
	f.SetCellValue("Sheet1", "H1", "Release Train")
	f.SetCellValue("Sheet1", "I1", "User")
	f.SetCellValue("Sheet1", "J1", "Created By")
	f.SetCellValue("Sheet1", "K1", "Status")

	// Percorrendo e colocando os Sheet1 do corpo da planilha (através da lista de products)
	for i, p := range list_remark.List {
		// Colocando os Sheet1 do corpo da planilha (lilha por lilha)
		f.SetCellValue("Sheet1", "A"+strconv.Itoa(i+2), p.ID)
		f.SetCellValue("Sheet1", "B"+strconv.Itoa(i+2), p.Subject_Name)
		f.SetCellValue("Sheet1", "C"+strconv.Itoa(i+2), p.Text)
		f.SetCellValue("Sheet1", "D"+strconv.Itoa(i+2), p.Date)
		f.SetCellValue("Sheet1", "E"+strconv.Itoa(i+2), p.Date_Return)
		f.SetCellValue("Sheet1", "F"+strconv.Itoa(i+2), p.Subject_Title)
		f.SetCellValue("Sheet1", "G"+strconv.Itoa(i+2), p.Client_Name)
		f.SetCellValue("Sheet1", "H"+strconv.Itoa(i+2), p.Release_Name)
		f.SetCellValue("Sheet1", "I"+strconv.Itoa(i+2), p.User_Name)
		f.SetCellValue("Sheet1", "J"+strconv.Itoa(i+2), p.CreatedBy_name)
		f.SetCellValue("Sheet1", "K"+strconv.Itoa(i+2), p.Status_Description)

		// Colocando formatação de estilo no corpo
		err = f.SetCellStyle("Sheet1", "A"+strconv.Itoa(i+2), "K"+strconv.Itoa(i+2), styleBody)
		// Verifica se teve algum erro ao colocar estilo
		if err != nil {
			return
		}
	}

	// Colocando tamanho na coluna B até F de 25
	err = f.SetColWidth("Sheet1", "B", "K", 25)
	// Verifica se teve algum erro ao colocar tamanho
	if err != nil {
		return
	}
	// Colocando formatação de estilo no cabeçalho
	err = f.SetCellStyle("Sheet1", "A1", "K1", styleHeader)
	// Verifica se teve algum erro ao colocar estilo
	if err != nil {
		return
	}

	// Salva a planilha Excel e coloca nome Products.xlsx
	if err := f.SaveAs("remarks.xlsx"); err != nil {
		fmt.Println(err)
	}
}
