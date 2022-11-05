resource "res_name" "res_id" {
  attribute1 = "value1"

  block1 "block1_label1" "block1_label2" "block1_label3" "block1_label4" {
    attribute21 = true
    attribute22 = 3
    block11 {
      attribute211 = ["A", "B"]

      attribute212 = -10

    }
  }

  attribute2 = "value2"
}
