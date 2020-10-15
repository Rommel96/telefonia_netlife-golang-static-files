var Assign = Vue.component("Assign", {
  template: `
  <template>
  <div class="container">
      <div class="row">
        <div>
          <div>
            <h2>Llenar los campos</h2>
            <form>
              <div class="form-group">
                <label for="id">Ingresa ID de la Diadema</label>
                <input
                  type="number"
                  v-model="headsetId"
                  name="headsetId"
                  class="form-control"
                />
              </div>
              <div class="form-group">
                <label htmlFor="agent">Ingresa nombre del Agente</label>
                <input
                  type="text"
                  v-model="agent"
                  name="agent"
                  class="form-control"
                />
              </div>
              <div class="form-group">
                <button @click.prevent="attachmetHeadset" class="btn btn-primary">
                  Asignar
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
      headsetId: 0,
      agent: "",
    };
  },
  methods: {
    async attachmetHeadset() {
      console.log(this.agent);
      const res = await fetch("/agent", {
        headers: {
          "Content-Type": "application/json",
        },
        method: "POST",
        body: JSON.stringify({
          headsetId: parseInt(this.headsetId),
          agent: this.agent,
        }),
      });
      const data = await res.json();
      console.log(data);
      this.$router.push("/report");
    },
  },
});
