package main

import (
	"net/http"
	"strings"

	"github.com/gouniverse/api"
	"github.com/gouniverse/cdn"
	"github.com/gouniverse/envenc"
	"github.com/gouniverse/hb"
	"github.com/gouniverse/utils"
	"github.com/gouniverse/webserver"
	"github.com/mingrammer/cfmt"
)

func keyAddModal() hb.TagInterface {
	buttonKeyAdd := hb.Button().
		Class("btn btn-primary").
		Text("Add key").
		Attr("v-on:click", "keyAdd()")
	buttonModalClose := hb.Button().
		Class("btn btn-secondary").
		Text("Close").
		Attr("v-on:click", "keyAddModalClose()")

	return hb.Div().
		ID("ModalKeyAdd").
		Class("modal").
		Child(hb.Div().
			Class("modal-dialog").
			Child(hb.Div().
				Class("modal-content").
				Child(hb.Div().
					Class("modal-header").
					Child(hb.H2().
						Style("margin: 0px;").
						Text("Add key"))).
				Child(hb.Div().
					Class("modal-body").
					Child(hb.Div().
						Class("form-group alert alert-danger").
						Attr("v-if", "keyAddForm.errorMessage").
						Child(hb.Span().Class("text-danger").Text("{{ keyAddForm.errorMessage }}")),
					).
					Child(hb.Div().
						Class("form-group").
						Child(hb.Label().
							Class("form-label").
							Text("Key")).
						Child(hb.Input().
							Class("form-control").
							Attr("type", "text").
							Attr("v-model", "keyAddForm.key"),
						),
					).
					Child(hb.Div().
						Class("form-group mt-3").
						Child(hb.Label().
							Class("form-label").
							Text("Value")).
						Child(hb.TextArea().
							Class("form-control").
							Attr("v-model", "keyAddForm.value"),
						),
					)).
				Child(
					hb.Div().
						Class("modal-footer justify-content-between").
						Child(buttonModalClose).
						Child(buttonKeyAdd),
				),
			),
		)
}

func keyRemoveModal() hb.TagInterface {
	buttonKeyRemove := hb.Button().
		Class("btn btn-danger").
		Text("Remove").
		Attr("v-on:click", "keyRemove()")
	buttonModalClose := hb.Button().
		Class("btn btn-secondary").
		Text("Close").
		Attr("v-on:click", "keyRemoveModalClose()")

	return hb.Div().
		ID("ModalKeyRemove").
		Class("modal").
		Child(hb.Div().
			Class("modal-dialog").
			Child(hb.Div().
				Class("modal-content").
				Child(hb.Div().
					Class("modal-header").
					Child(hb.H2().
						Style("margin: 0px;").
						Text("Remove key"))).
				Child(hb.Div().
					Class("modal-body").
					Child(hb.Div().
						Class("form-group").
						Attr("v-if", "keyRemoveForm.errorMessage").
						Child(hb.Span().Class("text-danger").Text("{{ keyRemoveForm.errorMessage }}")),
					).
					Child(hb.Div().
						Class("text-danger mb-3").
						Text("Are you sure you want to remove key '{{ keyRemoveForm.key }}'?"),
					).
					Child(hb.Div().
						Class("text-danger mb-3").
						Text("This action cannot be undone."),
					).
					Child(hb.Div().
						Class("form-group d-none").
						Child(hb.Label().
							Class("form-label").
							Text("Key")).
						Child(hb.Input().
							Class("form-control").
							Attr("type", "text").
							Attr("v-model", "keyRemoveForm.key").
							Attr("readonly", "readonly"),
						),
					)).
				Child(
					hb.Div().
						Class("modal-footer justify-content-between").
						Child(buttonModalClose).
						Child(buttonKeyRemove),
				),
			),
		)
}

func keyUpdateModal() hb.TagInterface {
	buttonKeyUpdate := hb.Button().
		Class("btn btn-primary").
		Text("Update").
		Attr("v-on:click", "keyUpdate(key)")
	buttonModalClose := hb.Button().
		Class("btn btn-secondary").
		Text("Close").
		Attr("v-on:click", "keyUpdateModalClose()")

	return hb.Div().
		ID("ModalKeyUpdate").
		Class("modal").
		Child(hb.Div().
			Class("modal-dialog").
			Child(hb.Div().
				Class("modal-content").
				Child(hb.Div().
					Class("modal-header").
					Child(hb.H2().
						Style("margin: 0px;").
						Text("Update key"))).
				Child(hb.Div().
					Class("modal-body").
					Child(hb.Div().
						Class("form-group alert alert-danger").
						Attr("v-if", "keyUpdateForm.errorMessage").
						Child(hb.Span().Class("text-danger").Text("{{ keyUpdateForm.errorMessage }}")),
					).
					Child(hb.Div().
						Class("form-group").
						Child(hb.Label().
							Class("form-label").
							Text("Key")).
						Child(hb.Input().
							Class("form-control").
							Style("background-color: #e9ecef;").
							Attr("type", "text").
							Attr("v-model", "keyUpdateForm.key").
							Attr("readonly", "readonly"),
						),
					).
					Child(hb.Div().
						Class("form-group mt-3").
						Child(hb.Label().
							Class("form-label").
							Text("Value")).
						Child(hb.TextArea().
							Class("form-control").
							Attr("v-model", "keyUpdateForm.value"),
						),
					)).
				Child(
					hb.Div().
						Class("modal-footer justify-content-between").
						Child(buttonModalClose).
						Child(buttonKeyUpdate),
				),
			),
		)
}

func keysPage() hb.TagInterface {
	buttonKeyNew := hb.Button().
		Class("btn btn-primary float-end").
		Text("➕ Add key").
		Attr("v-on:click", "keyAddModal()")

	title := hb.H1().
		Text("Keys").
		Child(buttonKeyNew)

	table := hb.Table().
		// Attr("v-if", "keys.length > 0").
		Class("table").
		Child(hb.Thead().
			Child(hb.TR().
				Child(hb.TH().Text("Key")).
				Child(hb.TH().Text("Value")).
				Child(hb.TH().Text("Actions")),
			),
		).
		Child(hb.Tbody().
			Child(hb.TR().
				Attr("v-for", "key in Object.keys(keys)").
				Child(hb.TD().Text("{{ key }}")).
				Child(hb.TD().Text("{{ keys[key].substr(0, 100) }}").Child(hb.Span().Attr("v-if", "keys[key].length > 100").Text("..."))).
				Child(hb.TD().
					Child(hb.Button().
						Class("btn btn-primary btn-sm me-2").
						Text("✏️ Update").
						Attr("v-on:click", "keyUpdateModalShow(key)")).
					Child(hb.Button().
						Class("btn btn-danger btn-sm").
						Text("❌ Remove").
						Attr("v-on:click", "keyRemoveModalShow(key)"))),
			),
		)

	return hb.Template().
		Attr("v-if", "pageKeysShow").
		Child(hb.Div().
			Class("container").
			Style("min-height:100vh").
			Style("margin-top:50px; margin-bottom:50px;").
			Child(title).
			Child(table).
			Child(keyUpdateModal()).
			Child(keyRemoveModal()).
			Child(keyAddModal()))
}

func loginPage() hb.TagInterface {
	buttonLogin := hb.Button().
		Class("btn btn-primary btn-lg mt-3 w-100").
		Text("Login").
		Attr("v-on:click", "login()")

	inputPassword := hb.Input().
		Class("form-control").
		Type(hb.TYPE_PASSWORD).
		Attr("v-model", "loginForm.password")

	inputVault := hb.Input().
		Class("form-control").
		Type(hb.TYPE_TEXT).
		Attr("v-model", "loginForm.vault")

	errorMessage := hb.Div().
		Class("alert alert-danger").
		Attr("v-if", "loginForm.errorMessage").
		Text("{{ loginForm.errorMessage }}")

	groupVault := hb.Div().
		Class("form-group mb-3").
		Child(hb.Label().
			Class("form-label").
			Text("Vault file path")).
		Child(inputVault)

	groupPassword := hb.Div().
		Class("form-group mb-3").
		Child(hb.Label().
			Class("form-label").
			Text("Vault file password")).
		Child(inputPassword)

	return hb.Template().
		Attr("v-if", "pageLoginShow").
		Child(
			hb.Div().
				Class("container").
				Style("margin-top:50px; margin-bottom:50px;").
				Style("height:100vh;").
				Style("display:flex;").
				Style("align-items:center;").
				Style("justify-content:center;").
				Child(hb.Div().
					Class("card").
					Style("width:100%;").
					Style("max-width:600px;").
					Child(hb.Div().
						Class("card-header").
						Child(hb.H1().
							Style("margin:0px;").
							Text("🔐 Login"))).
					Child(hb.Div().
						Class("card-body").
						Child(errorMessage).
						Child(groupVault).
						Child(groupPassword).
						Child(buttonLogin))))
}

func app() string {
	h := `
Vue.createApp({
	data() {
		return {
			pageKeysShow: false,
			pageLoginShow: true,
			keys: [],
			vaultPath: "",
			vaultPassword: "",
			keyAddForm: {
				key: "",
				value: "",
				errorMessage: "",
			},
			keyUpdateForm: {
				key: "",
				value: "",
				errorMessage: "",
			},
			keyRemoveForm: {
				key: "",
				errorMessage: "",
			},
			loginForm: {
				vault: "test.vault",
				password: "1234",
				errorMessage: "",
			},
		}
	},
	methods: {
		keyAdd() {
			$.post("?a=key-add", {
				vault: this.vaultPath,
				password: this.vaultPassword,
				key: this.keyAddForm.key,
				value: this.keyAddForm.value
			}).done((data) => {
				response = JSON.parse(data)

				if (response.status != "success") {
					this.keyAddForm.errorMessage = response.message;
					return;
				}
					
				this.keyAddForm.key = "";
				this.keyAddForm.value = "";
				this.keysList();
				this.keyAddModalClose();
			}).fail((data) => {
				this.keyAddForm.errorMessage = "Adding key failed";
			})
		},
		keyAddModal() {
			$("#ModalKeyAdd").modal().show();
		},
		keyAddModalClose() {
			$("#ModalKeyAdd").modal().hide();
		},
		keyUpdate() {
			$.post("?a=key-update", {
				vault: this.vaultPath,
				password: this.vaultPassword,
				key: this.keyUpdateForm.key,
				value: this.keyUpdateForm.value
			}).done((data) => {
				response = JSON.parse(data)
				if (response.status != "success") {
					this.keyUpdateForm.errorMessage = response.message;
					return;
				}
				this.keyUpdateForm.key = "";
				this.keyUpdateForm.value = "";
				this.keysList();
				this.keyUpdateModalClose();
			}).fail((data) => {
				this.keyUpdateForm.errorMessage = "Updating key failed";
			})
		},
		keyUpdateModalShow(key) {
			this.keyUpdateForm.key = key
			this.keyUpdateForm.value = this.keys[key]
			$("#ModalKeyUpdate").modal().show();
		},
		keyUpdateModalClose() {
			$("#ModalKeyUpdate").modal().hide();
		},
		keyRemoveModalShow(key) {
			this.keyRemoveForm.key = key
			$("#ModalKeyRemove").modal().show();
		},
		keyRemoveModalClose() {
			$("#ModalKeyRemove").modal().hide();
		},
		keyRemove() {
			$.post("?a=key-remove", {
				vault: this.vaultPath,
				password: this.vaultPassword,
				key: this.keyRemoveForm.key
			}).done((data) => {
				response = JSON.parse(data)

				if (response.status != "success") {
					this.keyRemoveForm.errorMessage = response.message;
					return;
				}

				this.keyRemoveForm.key = "";
				this.keyRemoveForm.errorMessage = "";
				this.keyRemoveModalClose();
				this.keysList()
			}).fail((data) => {
				this.keyRemoveForm.errorMessage = "Removing key failed";
			})
		},
		keysList() {
		   $.post("?a=keys", {
			   vault: this.vaultPath,
			   password: this.vaultPassword
		   }).done((data) => {
			   response = JSON.parse(data)

			   if (response.status != "success") {
				   this.loginForm.errorMessage = response.message;
				   this.keys = [];
				   this.pageKeysShow = false;
				   this.pageLoginShow = true;
				   return;
			   }

			   this.keys = response.data.keys
		   }).fail((data) => {
			   this.loginForm.errorMessage = "Listing keys failed";
		   })
		},
		login() {
			this.loginForm.errorMessage = ""

			$.post("?a=login", {
				vault: this.loginForm.vault,
				password: this.loginForm.password
			}).done((data) => {
				response = JSON.parse(data)

				if (response.status != "success") {			
					this.loginForm.errorMessage = response.message;
					return;
				}

				this.pageKeysShow = true;
				this.pageLoginShow = false;
				this.keys = response.data.keys;
				this.vaultPath = this.loginForm.vault;
				this.vaultPassword = this.loginForm.password;
			}).fail((data) => {
				this.loginForm.errorMessage = "Login failed";
			})
		}
	}
}).mount("#app")
	`
	return h
}

func webUi(w http.ResponseWriter, r *http.Request) string {
	action := utils.Req(r, "a", "")
	route := "/" + action

	router := map[string]func(w http.ResponseWriter, r *http.Request) string{
		"/key-add":    keyAdd,
		"/key-update": keyUpdate,
		"/key-remove": keyRemove,
		"/keys":       keys,
		"/login":      login,
		"/":           page,
	}

	if v, ok := router[route]; ok {
		return v(w, r)
	}

	return page(w, r)
}

func login(w http.ResponseWriter, r *http.Request) string {
	vault := strings.TrimSpace(utils.Req(r, "vault", ""))
	password := strings.TrimSpace(utils.Req(r, "password", ""))

	if vault == "" {
		return api.Error("Vault file is required").ToString()
	}

	if password == "" {
		return api.Error("Password is required").ToString()
	}

	if !utils.FileExists(vault) {
		return api.Error("Vault file does not exist").ToString()
	}

	data, err := envenc.KeyListFromFile(vault, password)

	if err != nil {
		return api.Error(err.Error()).ToString()
	}

	return api.SuccessWithData("Login successful", map[string]any{
		"keys": data,
	}).ToString()
}

func keyAdd(w http.ResponseWriter, r *http.Request) string {
	vault := strings.TrimSpace(utils.Req(r, "vault", ""))
	password := strings.TrimSpace(utils.Req(r, "password", ""))
	key := strings.TrimSpace(utils.Req(r, "key", ""))
	value := strings.TrimSpace(utils.Req(r, "value", ""))

	if vault == "" {
		return api.Error("Vault file is required").ToString()
	}

	if password == "" {
		return api.Error("Password is required").ToString()
	}

	if key == "" {
		return api.Error("Key is required").ToString()
	}

	if !utils.FileExists(vault) {
		return api.Error("Vault file does not exist").ToString()
	}

	exists, err := envenc.KeyExists(vault, password, key)

	if err != nil {
		return api.Error(err.Error()).ToString()
	}

	if exists {
		return api.Error("Key already exists").ToString()
	}

	if err := envenc.KeySet(vault, password, key, value); err != nil {
		return api.Error(err.Error()).ToString()
	}

	return api.Success("Key added successfully").ToString()
}

func keyRemove(w http.ResponseWriter, r *http.Request) string {
	vault := strings.TrimSpace(utils.Req(r, "vault", ""))
	password := strings.TrimSpace(utils.Req(r, "password", ""))
	key := strings.TrimSpace(utils.Req(r, "key", ""))

	if vault == "" {
		return api.Error("Vault file is required").ToString()
	}

	if password == "" {
		return api.Error("Password is required").ToString()
	}

	if key == "" {
		return api.Error("Key is required").ToString()
	}

	if !utils.FileExists(vault) {
		return api.Error("Vault file does not exist").ToString()
	}

	if err := envenc.KeyRemove(vault, password, key); err != nil {
		return api.Error(err.Error()).ToString()
	}

	return api.Success("Key removed successfully").ToString()
}

func keyUpdate(w http.ResponseWriter, r *http.Request) string {
	vault := strings.TrimSpace(utils.Req(r, "vault", ""))
	password := strings.TrimSpace(utils.Req(r, "password", ""))
	key := strings.TrimSpace(utils.Req(r, "key", ""))
	value := strings.TrimSpace(utils.Req(r, "value", ""))

	if vault == "" {
		return api.Error("Vault file is required").ToString()
	}

	if password == "" {
		return api.Error("Password is required").ToString()
	}

	if key == "" {
		return api.Error("Key is required").ToString()
	}

	if !utils.FileExists(vault) {
		return api.Error("Vault file does not exist").ToString()
	}

	if err := envenc.KeySet(vault, password, key, value); err != nil {
		return api.Error(err.Error()).ToString()
	}

	return api.Success("Key updated successfully").ToString()
}

func keys(w http.ResponseWriter, r *http.Request) string {
	vault := strings.TrimSpace(utils.Req(r, "vault", ""))
	password := strings.TrimSpace(utils.Req(r, "password", ""))

	if vault == "" {
		return api.Error("Vault file is required").ToString()
	}

	if password == "" {
		return api.Error("Password is required").ToString()
	}

	if !utils.FileExists(vault) {
		return api.Error("Vault file does not exist").ToString()
	}

	keys, err := envenc.KeyListFromFile(vault, password)

	if err != nil {
		return api.Error(err.Error()).ToString()
	}

	return api.SuccessWithData("Keys successful", map[string]any{
		"keys": keys,
	}).ToString()
}

func page(w http.ResponseWriter, r *http.Request) string {
	divApp := hb.Div().
		ID("app").
		Child(loginPage()).
		Child(keysPage())
	page := hb.NewWebpage()
	page.StyleURL(cdn.BootstrapUnitedCss_5_3_3())
	page.ScriptURL(cdn.BootstrapJs_5_3_3())
	page.ScriptURL(cdn.Htmx_2_0_0())
	page.ScriptURL(cdn.VueJs_3())
	page.ScriptURL(cdn.Jquery_3_7_1())
	page.Script(app())
	page.Child(divApp)
	return page.ToHTML()
}

// New UI in development
func main() {
	//envenc.NewCliV2().Run(args[0:])
	address := "127.0.0.1:38080"
	cfmt.Infoln("Listening on: http://" + address)

	s := webserver.New(address, func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(webUi(w, r)))
	})

	err := s.Start()

	if err != nil {
		panic(err)
	}
}
