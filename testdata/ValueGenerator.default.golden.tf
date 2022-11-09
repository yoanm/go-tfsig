sig {
  basic          = "basic_value"
  custom         = "custom.my_var"
  local          = local.my_local
  var            = var.my_var
  data           = data.my_data.my_property
  bool           = true
  number         = 1
  list_of_string = ["basic_value", "custom.my_var", local.my_local, var.my_var, data.my_data.my_property]
}
