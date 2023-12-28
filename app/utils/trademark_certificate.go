package utils

import (
	"github.com/icaksh/cripis/app/models"
	"github.com/jung-kurt/gofpdf"
)

func CreateCertificate(v *models.Trademark) (*gofpdf.Fpdf, error) {

	pdf := gofpdf.New("P", "mm", "A4", "")

	// Add a page to the PDF
	pdf.AddPage()

	// Set font
	pdf.SetFont("Arial", "", 12)

	// Define headers and data
	headers := []string{
		"Nomor dan tanggal permohonan",
		"Pemilik Merek Dagang",
		"Nama Alamat",
		"Nama Merek Dagang",
		"Tanggal diumumkan",
		"Tanggal Kadaluarsa",
	}

	data := []string{
		v.RegisterNumber + ", " + v.CreatedAt.Format("02 January 2006"),
		v.OwnerName,
		v.Address,
		v.TrademarkName,
		v.ApprovedAt.Time.Format("02 January 2006"),
		v.ExpiredAt.Time.Format("02 January 2006"),
	}

	// Set column widths
	colWidth := 80

	// Add headers in the first column
	for _, header := range headers {
		pdf.MultiCell(float64(colWidth), 10, header, "1", "C", false)
		pdf.Ln(-1)
	}

	pdf.SetX(float64(colWidth))
	for _, str := range data {
		pdf.MultiCell(float64(colWidth), 10, str, "1", "", false)
		pdf.Ln(-1)
	}
	return pdf, nil
}
