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
      </v-col>

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
                <!-- <v-text-field
                  label="ユーザー名"
                  placeholder="15文字以内"
                  outlined
                  dense
                  :rules="nameRules"
                ></v-text-field> -->
                <v-text-field
                  label="Eメールまたは携帯電話番号"
                  placeholder="mail@example.com"
                  outlined
                  dense
                  :rules="mailRules"
                ></v-text-field>
                <v-text-field
                  label="パスワード"
                  placeholder="8文字以上の半角英数記号"
                  outlined
                  dense
                  :rules="pwRules"
                ></v-text-field>
              </div>
              <!-- <div class="text-center">
                <v-btn block :disabled="!valid">ログイン</v-btn>
              </div> -->
              <div class="my-2">
                <v-btn block color="warning"> ログイン </v-btn>
              </div>
              <!-- <p class="signUp-border-top text-center mt-6 mb-0 pt-6">
                ソーシャルアカウントでログイン
              </p>
              <v-btn
                class="fill-width mt-6 text-capitalize caption mx-4"
                rounded
                color="#00ACEE"
                dark
                depressed
                height="48px"
                @click="submitTwitter"
              >
                <img
                  class="button-logo-img mr-4"
                  src="~/static/twitter.png"
                  style="height: 20px"
                />
                twitterでログイン
              </v-btn>
              <v-btn
                class="fill-width mt-6 text-capitalize caption mx-4 mb-6"
                rounded
                height="48px"
                outlined
                style="border-color: #979797"
                @click="submitGoogle"
              >
                <img
                  class="button-logo-img mr-4"
                  src="https://madeby.google.com/static/images/google_g_logo.svg"
                  style="height: 24px"
                />
                Googleでログイン
              </v-btn> -->
            </v-card-text>
          </v-form>
        </v-card>
      </div>
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
    };
  },
  methods: {
    validate() {
      this.$refs.form.validate();
    },
    submitTwitter() {
      // ツイッターログインの処理
    },
    submitGoogle() {
      // グーグルログインの処理
    },
  },
};
</script>
