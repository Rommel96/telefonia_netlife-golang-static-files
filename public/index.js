new Vue({
  el: "#login",
  data: {
    email: "",
    password: "",
  },
  methods: {
    async login() {
      const res = await fetch("/login", {
        headers: {
          "Content-Type": "application/json",
        },
        method: "POST",
        body: JSON.stringify({ email: this.email, password: this.password }),
      });
      const data = await res.json();
      console.log(data);
      this.email = "";
      this.password = "";
      if (data.email != "") window.location.href = "home.html";
    },
  },
});
