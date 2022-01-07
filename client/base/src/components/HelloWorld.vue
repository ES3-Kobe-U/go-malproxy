<template>
  <v-container>
    <v-row class="text-center">
      <v-col cols="12">
        <div class="d-flex flex-column justify-space-between align-center">
          <v-img
            :src="require('../assets/Amazon.co.jp_logo.svg.png')"
            class="my-3"
            contain
            max-height="150"
            max-width="200"
          />
        </div>
        <div>
          <v-card
            class="d-flex flex-column my-6 mx-auto"
            width="400"
            color="#fff"
          >
            <v-form ref="form" v-model="valid">
              <v-card-title class="d-flex justify-center pa-0 mt-6 mb-3"
                >Amazonアカウントでログイン</v-card-title
              >
              <v-card-text class="d-flex justify-center flex-column">
                <div class="mx-9">
                  <v-text-field
                    v-model="email"
                    label="Eメールまたは携帯電話番号"
                    placeholder="mail@example.com"
                    outlined
                    dense
                    :rules="mailRules"
                  ></v-text-field>
                  <v-text-field
                    v-model="password"
                    label="パスワード"
                    placeholder="8文字以上の半角英数記号"
                    outlined
                    dense
                    :rules="pwRules"
                  ></v-text-field>
                </div>
                <div class="my-2">
                  <v-btn block color="orange lighten-2"> ログイン </v-btn>
                </div>
              </v-card-text>
            </v-form>
          </v-card>
        </div>
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
export default {
  name: "HelloWorld",
  data() {
    return {
      valid: false,
      nameRules: [
        (v) => !!v || "user name is required",
        (v) => (v && v.length <= 15) || "最大15文字です。",
      ],
      mailRules: [
        (v) => !!v || "mail is required",
        (v) => /.+@.+\..+/.test(v) || "E-mail must be valid",
      ],
      pwRules: [(v) => !!v || "password is required"],

      email: "",
      password: "",
    };
  },
  methods: {
    validate() {
      this.$refs.form.validate();
    },
    go() {},
  },
};
</script>
