package main

import (
    "fmt"
    "github.com/xuri/excelize/v2"
    _ "image/gif"
    _ "image/jpeg"
    _ "image/png"
)

// create spreadsheet
func main_create_spreadsheet() {
    f := excelize.NewFile()
    //     create a new sheet
    index := f.NewSheet("Sheet2")
    //     set value of a cell
    f.SetCellValue("Sheet2", "A2", "Hello word.")
    f.SetCellValue("Sheet1", "B2", 100)
    // set active sheet of the workbook
    f.SetActiveSheet(index)
    // save spreadsheet by the given path
    if err := f.SaveAs("Book1.xlsx");
        err != nil {
        fmt.Println(err)
    }
}

// reading spreadsheet
func main_reading_spreadsheet() {
    f, err := excelize.OpenFile("Book1.xlsx")
    if err != nil {
        fmt.Println(err)
        return
    }
    // get value from cell by given worksheet name and axis.
    cell, err := f.GetCellValue("Sheet1", "B2")
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println(cell)

    fmt.Printf("============= %s ========\n", "get all the rows in the sheet1")
    // get all the rows in the sheet1
    rows, err := f.GetRows("Sheet1")
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println(rows)
    for _, row := range rows {
        for _, colCell := range row {
            fmt.Println(colCell, "\t")
        }
        fmt.Println()
    }
}

// add chart to spreadsheet file
func main_add_chart() {
    categories := map[string]string{
        "A2": "Small", "A3": "Normal", "A4": "Large",
        "B1": "Apple", "C1": "Orange", "D1": "Pear"}
    values := map[string]int{
        "B2": 2, "C2": 3, "D2": 3, "B3": 5, "C3": 2, "D3": 4, "B4": 6, "C4": 7, "D4": 8}
    f := excelize.NewFile()
    for k, v := range categories {
        f.SetCellValue("Sheet1", k, v)
    }
    for k, v := range values {
        f.SetCellValue("Sheet1", k, v)
    }
    if err := f.AddChart("Sheet1", "E1", `{
        "type": "col3DClustered",
        "series": [
        {
            "name": "Sheet1!$A$2",
            "categories": "Sheet1!$B$1:$D$1",
            "values": "Sheet1!$B$2:$D$2"
        },
        {
            "name": "Sheet1!$A$3",
            "categories": "Sheet1!$B$1:$D$1",
            "values": "Sheet1!$B$3:$D$3"
        },
        {
            "name": "Sheet1!$A$4",
            "categories": "Sheet1!$B$1:$D$1",
            "values": "Sheet1!$B$4:$D$4"
        }],
        "title":
        {
            "name": "Fruit 3D Clustered Column Chart"
        }
    }`); err != nil {
        fmt.Println(err)
        return
    }
    // save spreadsheet by given path.
    if err := f.SaveAs("Book1.xlsx"); err != nil {
        fmt.Println(err)
    }

}

// add picture to spreadsheet file
func main() {
    f, err := excelize.OpenFile("Book1.xlsx")
    if err != nil {
        fmt.Println(err)
        return
    }
    //     insert a picture
    if err := f.AddPicture("Sheet1", "A2", "image/image.png", ""); err != nil {
        fmt.Println(err)
    }
    // insert a picture to worksheet with scaling
    if err := f.AddPicture("Sheet1", "D2", "image/image.jpg", `{"x_scale":0.5,"y_scale":0.5}`); err != nil {
        fmt.Println(err)
    }
    // insert a picture offset in teh cell with pringing support
    if err := f.AddPicture("Sheet1", "H2", "image/image.gif", `{
        "o_offset":15,
        "y_offset":10,
        "print_obj": true,
        "lock_aspect_ratio": false,
        "locked": false }`); err != nil {
        fmt.Println(err)
    }
    // save the spreadsheet with the origin path
    if err := f.Save(); err != nil {
        fmt.Println(err)
    }
}
