package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["app-rest-inventory/controllers:BillsController"] = append(beego.GlobalControllerRouter["app-rest-inventory/controllers:BillsController"],
		beego.ControllerComments{
			Method: "CreateBill",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["app-rest-inventory/controllers:BillsController"] = append(beego.GlobalControllerRouter["app-rest-inventory/controllers:BillsController"],
		beego.ControllerComments{
			Method: "GetBills",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(
				param.New("from"),
				param.New("to"),
			),
			Params: nil})

	beego.GlobalControllerRouter["app-rest-inventory/controllers:BillsController"] = append(beego.GlobalControllerRouter["app-rest-inventory/controllers:BillsController"],
		beego.ControllerComments{
			Method: "GetBill",
			Router: `/:bill_id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(
				param.New("bill_id", param.IsRequired, param.InPath),
			),
			Params: nil})

	beego.GlobalControllerRouter["app-rest-inventory/controllers:BillsController"] = append(beego.GlobalControllerRouter["app-rest-inventory/controllers:BillsController"],
		beego.ControllerComments{
			Method: "UpdateDiscount",
			Router: `/:bill_id`,
			AllowHTTPMethods: []string{"patch"},
			MethodParams: param.Make(
				param.New("bill_id", param.IsRequired, param.InPath),
			),
			Params: nil})

	beego.GlobalControllerRouter["app-rest-inventory/controllers:BillsController"] = append(beego.GlobalControllerRouter["app-rest-inventory/controllers:BillsController"],
		beego.ControllerComments{
			Method: "DeleteBill",
			Router: `/:bill_id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(
				param.New("bill_id", param.IsRequired, param.InPath),
			),
			Params: nil})

	beego.GlobalControllerRouter["app-rest-inventory/controllers:BillsController"] = append(beego.GlobalControllerRouter["app-rest-inventory/controllers:BillsController"],
		beego.ControllerComments{
			Method: "AddSale",
			Router: `/:bill_id/sales/:sale_id`,
			AllowHTTPMethods: []string{"patch"},
			MethodParams: param.Make(
				param.New("bill_id", param.IsRequired, param.InPath),
				param.New("sale_id", param.IsRequired, param.InPath),
			),
			Params: nil})

	beego.GlobalControllerRouter["app-rest-inventory/controllers:BillsController"] = append(beego.GlobalControllerRouter["app-rest-inventory/controllers:BillsController"],
		beego.ControllerComments{
			Method: "RemoveSale",
			Router: `/:bill_id/sales/:sale_id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(
				param.New("bill_id", param.IsRequired, param.InPath),
				param.New("sale_id", param.IsRequired, param.InPath),
			),
			Params: nil})

	beego.GlobalControllerRouter["app-rest-inventory/controllers:CateringsController"] = append(beego.GlobalControllerRouter["app-rest-inventory/controllers:CateringsController"],
		beego.ControllerComments{
			Method: "CreateCatering",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["app-rest-inventory/controllers:CateringsController"] = append(beego.GlobalControllerRouter["app-rest-inventory/controllers:CateringsController"],
		beego.ControllerComments{
			Method: "GetCaterings",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(
				param.New("from"),
				param.New("to"),
			),
			Params: nil})

	beego.GlobalControllerRouter["app-rest-inventory/controllers:CateringsController"] = append(beego.GlobalControllerRouter["app-rest-inventory/controllers:CateringsController"],
		beego.ControllerComments{
			Method: "GetCatering",
			Router: `/:catering_id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(
				param.New("catering_id", param.IsRequired, param.InPath),
			),
			Params: nil})

	beego.GlobalControllerRouter["app-rest-inventory/controllers:CateringsController"] = append(beego.GlobalControllerRouter["app-rest-inventory/controllers:CateringsController"],
		beego.ControllerComments{
			Method: "UpdateCatering",
			Router: `/:catering_id`,
			AllowHTTPMethods: []string{"patch"},
			MethodParams: param.Make(
				param.New("catering_id", param.IsRequired, param.InPath),
			),
			Params: nil})

	beego.GlobalControllerRouter["app-rest-inventory/controllers:CustomersController"] = append(beego.GlobalControllerRouter["app-rest-inventory/controllers:CustomersController"],
		beego.ControllerComments{
			Method: "CreateCustomer",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["app-rest-inventory/controllers:CustomersController"] = append(beego.GlobalControllerRouter["app-rest-inventory/controllers:CustomersController"],
		beego.ControllerComments{
			Method: "GetCustomer",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["app-rest-inventory/controllers:HeadquartersController"] = append(beego.GlobalControllerRouter["app-rest-inventory/controllers:HeadquartersController"],
		beego.ControllerComments{
			Method: "CreateHeadquarter",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["app-rest-inventory/controllers:HeadquartersController"] = append(beego.GlobalControllerRouter["app-rest-inventory/controllers:HeadquartersController"],
		beego.ControllerComments{
			Method: "GetHeadquarters",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["app-rest-inventory/controllers:HeadquartersController"] = append(beego.GlobalControllerRouter["app-rest-inventory/controllers:HeadquartersController"],
		beego.ControllerComments{
			Method: "GetHeadquarter",
			Router: `/:headquarter_id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(
				param.New("headquarter_id", param.IsRequired, param.InPath),
			),
			Params: nil})

	beego.GlobalControllerRouter["app-rest-inventory/controllers:HeadquartersController"] = append(beego.GlobalControllerRouter["app-rest-inventory/controllers:HeadquartersController"],
		beego.ControllerComments{
			Method: "UpdateHeadquarter",
			Router: `/:headquarter_id`,
			AllowHTTPMethods: []string{"patch"},
			MethodParams: param.Make(
				param.New("headquarter_id", param.IsRequired, param.InPath),
			),
			Params: nil})

	beego.GlobalControllerRouter["app-rest-inventory/controllers:HeadquartersController"] = append(beego.GlobalControllerRouter["app-rest-inventory/controllers:HeadquartersController"],
		beego.ControllerComments{
			Method: "DeleteHeadquarter",
			Router: `/:headquarter_id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(
				param.New("headquarter_id", param.IsRequired, param.InPath),
			),
			Params: nil})

	beego.GlobalControllerRouter["app-rest-inventory/controllers:HeadquartersController"] = append(beego.GlobalControllerRouter["app-rest-inventory/controllers:HeadquartersController"],
		beego.ControllerComments{
			Method: "AddProduct",
			Router: `/:headquarter_id/products`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(
				param.New("headquarter_id", param.IsRequired, param.InPath),
			),
			Params: nil})

	beego.GlobalControllerRouter["app-rest-inventory/controllers:HeadquartersController"] = append(beego.GlobalControllerRouter["app-rest-inventory/controllers:HeadquartersController"],
		beego.ControllerComments{
			Method: "GetProducts",
			Router: `/:headquarter_id/products`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(
				param.New("headquarter_id", param.IsRequired, param.InPath),
				param.New("name"),
				param.New("brand"),
				param.New("color"),
			),
			Params: nil})

	beego.GlobalControllerRouter["app-rest-inventory/controllers:HeadquartersController"] = append(beego.GlobalControllerRouter["app-rest-inventory/controllers:HeadquartersController"],
		beego.ControllerComments{
			Method: "RemoveProducts",
			Router: `/:headquarter_id/products`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(
				param.New("headquarter_id", param.IsRequired, param.InPath),
			),
			Params: nil})

	beego.GlobalControllerRouter["app-rest-inventory/controllers:HeadquartersController"] = append(beego.GlobalControllerRouter["app-rest-inventory/controllers:HeadquartersController"],
		beego.ControllerComments{
			Method: "UpdateProduct",
			Router: `/:headquarter_id/products/:product_id`,
			AllowHTTPMethods: []string{"patch"},
			MethodParams: param.Make(
				param.New("headquarter_id", param.IsRequired, param.InPath),
				param.New("product_id", param.IsRequired, param.InPath),
			),
			Params: nil})

	beego.GlobalControllerRouter["app-rest-inventory/controllers:HeadquartersController"] = append(beego.GlobalControllerRouter["app-rest-inventory/controllers:HeadquartersController"],
		beego.ControllerComments{
			Method: "GetProduct",
			Router: `/:headquarter_id/products/:product_id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(
				param.New("headquarter_id", param.IsRequired, param.InPath),
				param.New("product_id", param.IsRequired, param.InPath),
			),
			Params: nil})

	beego.GlobalControllerRouter["app-rest-inventory/controllers:ProductsController"] = append(beego.GlobalControllerRouter["app-rest-inventory/controllers:ProductsController"],
		beego.ControllerComments{
			Method: "CreateProduct",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["app-rest-inventory/controllers:ProductsController"] = append(beego.GlobalControllerRouter["app-rest-inventory/controllers:ProductsController"],
		beego.ControllerComments{
			Method: "GetProducts",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(
				param.New("name"),
				param.New("brand"),
				param.New("color"),
			),
			Params: nil})

	beego.GlobalControllerRouter["app-rest-inventory/controllers:ProductsController"] = append(beego.GlobalControllerRouter["app-rest-inventory/controllers:ProductsController"],
		beego.ControllerComments{
			Method: "DeleteCatering",
			Router: `/:catering_id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(
				param.New("catering_id", param.IsRequired, param.InPath),
			),
			Params: nil})

	beego.GlobalControllerRouter["app-rest-inventory/controllers:ProductsController"] = append(beego.GlobalControllerRouter["app-rest-inventory/controllers:ProductsController"],
		beego.ControllerComments{
			Method: "GetProduct",
			Router: `/:product_id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(
				param.New("product_id", param.IsRequired, param.InPath),
			),
			Params: nil})

	beego.GlobalControllerRouter["app-rest-inventory/controllers:ProductsController"] = append(beego.GlobalControllerRouter["app-rest-inventory/controllers:ProductsController"],
		beego.ControllerComments{
			Method: "UpdateProduct",
			Router: `/:product_id`,
			AllowHTTPMethods: []string{"patch"},
			MethodParams: param.Make(
				param.New("product_id", param.IsRequired, param.InPath),
			),
			Params: nil})

	beego.GlobalControllerRouter["app-rest-inventory/controllers:ProductsController"] = append(beego.GlobalControllerRouter["app-rest-inventory/controllers:ProductsController"],
		beego.ControllerComments{
			Method: "DeleteProduct",
			Router: `/:product_id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(
				param.New("product_id", param.IsRequired, param.InPath),
			),
			Params: nil})

	beego.GlobalControllerRouter["app-rest-inventory/controllers:ProductsController"] = append(beego.GlobalControllerRouter["app-rest-inventory/controllers:ProductsController"],
		beego.ControllerComments{
			Method: "GetBrands",
			Router: `/brands`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["app-rest-inventory/controllers:ProvidersController"] = append(beego.GlobalControllerRouter["app-rest-inventory/controllers:ProvidersController"],
		beego.ControllerComments{
			Method: "CreateProvider",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["app-rest-inventory/controllers:ProvidersController"] = append(beego.GlobalControllerRouter["app-rest-inventory/controllers:ProvidersController"],
		beego.ControllerComments{
			Method: "GetProviders",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["app-rest-inventory/controllers:ProvidersController"] = append(beego.GlobalControllerRouter["app-rest-inventory/controllers:ProvidersController"],
		beego.ControllerComments{
			Method: "GetProvider",
			Router: `/:provider_id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(
				param.New("provider_id", param.IsRequired, param.InPath),
			),
			Params: nil})

	beego.GlobalControllerRouter["app-rest-inventory/controllers:ProvidersController"] = append(beego.GlobalControllerRouter["app-rest-inventory/controllers:ProvidersController"],
		beego.ControllerComments{
			Method: "UpdateProvider",
			Router: `/:provider_id`,
			AllowHTTPMethods: []string{"patch"},
			MethodParams: param.Make(
				param.New("provider_id", param.IsRequired, param.InPath),
			),
			Params: nil})

	beego.GlobalControllerRouter["app-rest-inventory/controllers:ProvidersController"] = append(beego.GlobalControllerRouter["app-rest-inventory/controllers:ProvidersController"],
		beego.ControllerComments{
			Method: "DeleteProvider",
			Router: `/:provider_id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(
				param.New("provider_id", param.IsRequired, param.InPath),
			),
			Params: nil})

	beego.GlobalControllerRouter["app-rest-inventory/controllers:UsersController"] = append(beego.GlobalControllerRouter["app-rest-inventory/controllers:UsersController"],
		beego.ControllerComments{
			Method: "CreateUser",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["app-rest-inventory/controllers:UsersController"] = append(beego.GlobalControllerRouter["app-rest-inventory/controllers:UsersController"],
		beego.ControllerComments{
			Method: "GetUsers",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["app-rest-inventory/controllers:UsersController"] = append(beego.GlobalControllerRouter["app-rest-inventory/controllers:UsersController"],
		beego.ControllerComments{
			Method: "GetUser",
			Router: `/:user_id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(
				param.New("user_id", param.IsRequired, param.InPath),
			),
			Params: nil})

	beego.GlobalControllerRouter["app-rest-inventory/controllers:UsersController"] = append(beego.GlobalControllerRouter["app-rest-inventory/controllers:UsersController"],
		beego.ControllerComments{
			Method: "UpdateUser",
			Router: `/:user_id`,
			AllowHTTPMethods: []string{"patch"},
			MethodParams: param.Make(
				param.New("user_id", param.IsRequired, param.InPath),
			),
			Params: nil})

	beego.GlobalControllerRouter["app-rest-inventory/controllers:UsersController"] = append(beego.GlobalControllerRouter["app-rest-inventory/controllers:UsersController"],
		beego.ControllerComments{
			Method: "DeleteUser",
			Router: `/:user_id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(
				param.New("user_id", param.IsRequired, param.InPath),
			),
			Params: nil})

}
