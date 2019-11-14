resource "sdm_node" "relay" {
    relay {
        name = "blargh"
    }
}

resource "sdm_node" "gateway" {
    gateway {
        name = "gateway"
        listen_address = "sdmlocal:21222"
        bind_address = "0.0.0.0:0"
    }
}

resource "sdm_role" "compositeroletest" {
    name = "compositeroletest-BLARGH"
    composite = true
}

resource "sdm_role" "roletest" {
    name = "roletest-BLARGH"
    composite = false
}