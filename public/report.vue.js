var Report = Vue.component("Report", {
  template: `
  <template v-if="!report" >
  ----
  </template >
  <template v-else>
  <table class="table table-bordered">
    <thead style="text-align: center;">
      <tr>
      <th>Serie</th>
      <th>Marca</th>
      <th>Agente</th>
      </tr>
    </thead>
    <tbody >
      <tr v-for="entry in report">
        <td>
          {{entry.serial}}
        </td>
        <td>
          {{entry.marca}}
        </td>
        <td>
          {{entry.agente}}
        </td>
      </tr>
    </tbody>
  </table>
  </template >
  `,
  data: function () {
    return {
      report: [],
    };
  },
  created() {
    this.getReport();
  },
  methods: {
    getReport: async function () {
      const res = await fetch("/report", {
        headers: {
          "Content-Type": "application/json",
        },
        method: "POST",
      });
      const data = await res.json();
      console.log(data);
      this.report = data;
    },
  },
});
