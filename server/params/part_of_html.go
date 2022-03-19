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

var RakutenReplaceNo1 = `<div class="spacer--xFAdr  block--2PK_L  
        
        
        
        padding-bottom-small--sgTI2
       
      
      
      
      border-bottom-gray--OXbtM
      
     white--3LZcf"><div class="spacer--xFAdr  block--2PK_L  
        padding-left-xlarge--2d9GV
        padding-right-xlarge--LeKQw
        
        padding-bottom-xxsmall--14_zk
        "><div class="container--IckCk type-3--1HhJJ"><span class="clickable-span--15gHd"><div class="icon--2tjYQ icon-left--3FsRA"><div class="text-display--1Iony type-icon--3g0D- size-custom-medium--3iEUT align-left--1hi1x color-information-icon--3Z3gZ  layout-inline--1ajCj"><div class="icon--2sY_j common-info-filled--auWfJ"></div></div></div><div class="text--2TE80"><div class="text-display--1Iony type-body--1W5uC size-small--sv6IW align-left--1hi1x color-gray-darker--1SJFG line-height-medium--2-H3z layout-inline--1ajCj">ウクライナ 緊急支援募金のお知らせ</div></div><div class="icon--2tjYQ icon-right--2F1nI"><div class="text-display--1Iony type-icon--3g0D- size-custom-small--2Y-pv align-right--2ACTn color-gray--1TFBo  layout-inline--1ajCj"><div class="icon--2sY_j common-chevron-right--VZMgW"></div></div></div></span></div></div></div>`

var RakutenReplaceNo2 = `<div irc="CommonHeaderMall" data-url="https://www.rakuten.co.jp" data-settings="[
      {
  &quot;tracker&quot;: {
    &quot;params&quot;: {
      &quot;accountId&quot;: 1,
      &quot;serviceId&quot;: 1,
      &quot;pageLayout&quot;: &quot;pc&quot;,
      &quot;pageType&quot;: &quot;top&quot;
    }
  },

  &quot;showSearchBar&quot;: true,
  &quot;showMemberInfoSummary&quot;: false,
  &quot;showSpu&quot;: false,
  &quot;showCartModal&quot;: true,
  &quot;customLogoImageUrl&quot;: &quot;https://r.r10s.jp/com/img/thumb/logo/logo_rakuten_25th.svg&quot;,
      &quot;links&quot;: {
        &quot;top&quot;: &quot;https://corp.rakuten.co.jp/event/anniversary25th/?scid=wi_ich_r25_pc_top_header_25th_logo_v1&quot;
      },
  &quot;withBorder&quot;: false,
  &quot;suggestionUrl&quot; : &quot;https://rdc-api-catalog-gateway-api.rakuten.co.jp/SUI/autocomplete/pc&quot;,
  &quot;useTBasketDomain&quot;: true,
  &quot;api&quot;: {
      &quot;cartApiSid&quot;: 1010,
      &quot;notificationLocId&quot;: 25,
      &quot;url&quot;:&quot;https://api-ichiba-gateway.rakuten.co.jp/graphql-common-bff/graphql&quot;,
      &quot;apikey&quot;: &quot;59093b15781a092c9573ea7032016ddb&quot;,
      &quot;clientId&quot;: &quot;top&quot;,
      &quot;spuViewType&quot;: &quot;top&quot;,
      &quot;spuSource&quot;: &quot;pc&quot;,
      &quot;spuEncoding&quot;: &quot;UTF-8&quot;,
      &quot;spuAcc&quot;: 1,
      &quot;spuAid&quot;: 1
    }
}

    ]">`
