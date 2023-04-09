package pdf

import (
	"fmt"
	"github.com/johnfercher/maroto/pkg/color"
	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/pdf"
	"github.com/johnfercher/maroto/pkg/props"
	"log"
	"main/model"
	"time"
)

func CreatePDF(order model.Order) error {
	m := pdf.NewMaroto(consts.Portrait, consts.A4)
	m.SetPageMargins(20, 10, 20)
	buildHeading(m, order.OrderID)
	buildOrder(m, order.Books)
	buildTotalAmount(m, order.TotalPrice)

	m.SetAuthor("hf42", true)
	m.SetCreator("david", true)
	m.SetTitle("Invoice", true)
	m.SetSubject("Invoice", true)
	m.SetCreationDate(time.Now())

	return m.OutputFileAndClose("invoice.pdf")
}

func buildHeading(m pdf.Maroto, orderID string) {
	m.RegisterHeader(func() {
		m.Row(50, func() {
			m.Col(12, func() {
				err := m.FileImage("logo/logo.png", props.Rect{
					Center:  true,
					Percent: 100,
				})

				if err != nil {
					log.Printf("error with file image: %s\n", err)
				}

			})
		})
	})

	m.Row(10, func() {
		m.Col(12, func() {
			m.Text("Invoice number: "+orderID, props.Text{
				Top:   3,
				Style: consts.Bold,
				Align: consts.Center,
				Color: getDarkPurpleColor(),
			})
		})
	})
}

func getDarkPurpleColor() color.Color {
	return color.Color{
		Red:   88,
		Green: 80,
		Blue:  99,
	}
}

func buildOrder(m pdf.Maroto, books []model.Book) {
	m.SetBackgroundColor(getTealColor())

	tableHeadings := []string{"Title", "Quantity", "Price", "Total amount"}

	m.Row(10, func() {
		m.Col(12, func() {
			m.Text("Books", props.Text{
				Top:    2,
				Size:   13,
				Color:  color.NewWhite(),
				Family: consts.Courier,
				Style:  consts.Bold,
				Align:  consts.Center,
			})
		})
	})

	contents := make([][]string, len(books))

	for k, book := range books {
		contents[k] = []string{
			book.Title,
			fmt.Sprintf("%d", book.Quantity),
			fmt.Sprintf("%.2f€", book.Price),
			fmt.Sprintf("%.2f€", book.TotalPrice),
		}
	}

	m.SetBackgroundColor(color.NewWhite())

	lightPurpleColor := getLightPurpleColor()
	m.TableList(tableHeadings, contents, props.TableList{
		HeaderProp: props.TableListContent{
			Size:      12,
			GridSizes: []uint{5, 2, 2, 3},
		},
		ContentProp: props.TableListContent{
			Size:      12,
			GridSizes: []uint{5, 2, 2, 3},
		},
		Align:                  consts.Left,
		AlternatedBackground:   &lightPurpleColor,
		HeaderContentSpace:     1,
		VerticalContentPadding: 1,
		Line:                   false,
	})
}

func buildTotalAmount(m pdf.Maroto, amount float32) {
	m.Row(20, func() {
		m.Col(1, func() {
			m.Text("Total: ", props.Text{
				Top:   10,
				Style: consts.Bold,
				Size:  15,
				Align: consts.Right,
			})
		})
		m.Col(2, func() {
			m.Text(fmt.Sprintf("%.2f€", amount), props.Text{
				Top:   15,
				Style: consts.Bold,
				Size:  15,
				Align: consts.Center,
			})
		})
	})
}

func getLightPurpleColor() color.Color {
	return color.Color{
		Red:   210,
		Green: 200,
		Blue:  230,
	}
}

func getTealColor() color.Color {
	return color.Color{
		Red:   3,
		Green: 166,
		Blue:  166,
	}
}
