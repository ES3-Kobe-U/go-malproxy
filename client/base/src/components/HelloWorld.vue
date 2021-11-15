<template>
  <div
    class="bg-body d-flex flex-column justify-center align-center min-vh-100"
  >
    <div class="sign-up-card">
      <div class="sign-up-card-container">
        <div class="text-center mb-10">
          <h3 class="mb-3">Authentication</h3>
          <h5 class="text-sm font-600 grey--text text--darken-4">
            Authentication will be performed to confirm your identity.
            <br />Enter the name and login password of your PC's administrator
            account.
          </h5>
        </div>
        <div class="mb-4">
          <p class="text-14 mb-1">Username</p>
          <v-text-field
            placeholder="Input your account name"
            v-model="email"
            outlined
            dense
            hide-details=""
            class="mb-4"
          ></v-text-field>
        </div>
        <div class="mb-4">
          <p class="text-14 mb-1">Password</p>
          <v-text-field
            type="password"
            v-model="password"
            placeholder="Input your login pass"
            outlined
            dense
            hide-details=""
            class="mb-4"
          ></v-text-field>
        </div>

        <div class="mb-4">
          <v-btn
            @click="signin"
            block
            color="primary"
            class="text-capitalize font-600"
          >
            Start authentication
          </v-btn>
        </div>
        <div
          class="d-flex align-center justify-center w-200 mx-auto mb-4"
        ></div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: "HelloWorld",

  data: () => ({
    errorMessages: "",
    name: null,
    address: null,
    city: null,
    state: null,
    zip: null,
    country: null,
    formHasErrors: false,
  }),

  computed: {
    form() {
      return {
        name: this.name,
        address: this.address,
        city: this.city,
        state: this.state,
        zip: this.zip,
        country: this.country,
      };
    },
  },

  watch: {
    name() {
      this.errorMessages = "";
    },
  },

  methods: {
    addressCheck() {
      this.errorMessages =
        this.address && !this.name ? `Hey! I'm required` : "";

      return true;
    },
    resetForm() {
      this.errorMessages = [];
      this.formHasErrors = false;

      Object.keys(this.form).forEach((f) => {
        this.$refs[f].reset();
      });
    },
    submit() {
      this.formHasErrors = false;

      Object.keys(this.form).forEach((f) => {
        if (!this.form[f]) this.formHasErrors = true;

        this.$refs[f].validate(true);
      });
    },
  },
};
</script>
<style lang="scss" scoped>
.sign-up-card {
  width: 500px;
  overflow: hidden;
  background-color: #fff;
  border-radius: 8px;
  margin: 2rem auto;
  box-shadow: rgb(3 0 71 / 9%) 0px 8px 45px;
  @media (max-width: 500px) {
    width: 100%;
  }
  .sign-up-card-container {
    padding: 3rem 3.75rem 0px;
    @media (max-width: 500px) {
      padding: 3rem 1rem 0px;
    }
  }
}
</style>
