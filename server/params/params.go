package params

var RakutenLoginCode = `<div
    irc="MembershipHeader"
    data-show-new="true"
    data-tracker='{
  "params": {
      "accountId": 1,
      "serviceId": 1,
      "pageLayout": "pc",
      "pageType": "top"
    }
}
'
    ></div>`

var ReplaceRakutenLoginCode = `<div class="header--RqsDU anonymous--2RuTn">
        <ul class="main-menu--33xG4">
          <li class="section--1itbn">
            <div
              class="text-display--1Iony type-body--1W5uC size-x-large--20opE align-left--1hi1x color-gray-darker--1SJFG  layout-inline--1ajCj">
              ようこそ楽天市場へ</div>
          </li>
          <li>
            <div
              class="text-display--1Iony type-body--1W5uC size-x-large--20opE align-left--1hi1x color-gray-dark--2N4Oj  layout-inline--1ajCj">
              会員登録で楽天ポイントが貯まる、使える。</div>
          </li>
        </ul>
        <ul class="side-menu--37357">
          <li class="section--1itbn"><a
              class="button--3SNaj size-s--KzHQM size-s-padding--AtFL_ border-radius--1ip29 no-padding--3mzqd type-link--8tP4V type-link-icon--2KEwc variant-crimson--3DbX7"
              aria-label="楽天会員登録(無料)"
              href="https://rd.rakuten.co.jp/s/?R2=https%3A%2F%2Fgrp01.id.rakuten.co.jp%2Frms%2Fnid%2Fregistfwd%3Fservice_id%3Dtop&amp;D2=3.8611.68708.907372.32326946&amp;C3=733f3c7a572082b53fe39af9150e9b3503e19bf2"
              target="_self" aria-disabled="false" aria-pressed="false" tabindex="0">
              <div class="icon--2sY_j size-m--23dCu color-crimson--2DVXa rex-user-outline--3_CEJ"></div><span
                class="text--26ZD7 text-no-margin-right--3R22- text--xKo_r">楽天会員登録(無料)</span>
            </a></li>
          <li class="section--1itbn"><button
              class="button--3SNaj size-xs--2qQUS size-xs-padding--pvhl0 border-radius--1ip29 type-primary--3cgWx"
              aria-label="ログイン" type="button" onclick="location.href='/rakuten-login'">
              <div class="icon--2sY_j size-s--By3wJ color-white--fjaFR rex-login--12L7t"></div><span
                class="text--26ZD7 text-no-margin-right--3R22- text--76coE">ログイン</span>
            </button></li>
        </ul>
      </div>`
