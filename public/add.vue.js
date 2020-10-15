var Add = Vue.component("Add", {
  template: `
  <template>
  <div class="container">
      <div class="row">
        <div>
          <div>
            <h2>Llenar los campos</h2>
            <form>
              <div class="form-group">
                <label for="serial">Ingresa numero de serie</label>
                <input
                  type="text"
                  v-model="serial"
                  name="serial"
                  class="form-control"
                />
              </div>
              <div class="form-group">
                <label htmlFor="marca">Ingresa la marca</label>
                <input
                  type="text"
                  v-model="marca"
                  name="marca"
                  class="form-control"
                />
              </div>
              <div class="form-group">
                <button @click.prevent="insertHeadset" class="btn btn-primary">
                  Agregar
                </button>
              </div>
            </form>
          </div>
        </div>
      </div>
    </div>
  </template>
  `,
  data: function () {
    return {
      serial: "",
      marca: "",
    };
  },
  methods: {
    async insertHeadset() {
      const res = await fetch("/headset", {
        headers: {
          "Content-Type": "application/json",
        },
        method: "POST",
        body: JSON.stringify({ serial: this.serial, marca: this.marca }),
      });
      const data = await res.json();
      console.log(data);
      this.$router.push("/report");
    },
  },
});
