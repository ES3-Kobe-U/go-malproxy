var __px = window.__px || {};
__px.campaigns = __px.campaigns || [];
__px.campaigns.push({
    "nest": {
        "conditions": [{
            "patternName": "target__root__nonlogin",
            "patternId": "target__49869__176320",
            "actions": {
                "replace": [{
                    "source": "/ichibatop/com/inc/home/20080930/beta/tz/liquid/grouplist/20210416/pitari_nonlogin.html",
                    "target": "group_banner_pitari"
                }]
            },
            "match": {
                "__pitari_non_login": 0
            }
        }, {
            "nest": {
                "conditions": [{
                    "patternName": "target__target_card_scv__non_bank",
                    "patternId": "target__49870__176321",
                    "actions": {
                        "replace": [{
                            "source": "/ichibatop/com/inc/home/20080930/beta/tz/liquid/grouplist/20210416/pitari_non_bank.html",
                            "target": "group_banner_pitari"
                        }]
                    },
                    "match": {
                        "cu_bankscv": 0
                    }
                }, {
                    "patternName": "target__target_card_scv__non_sec",
                    "patternId": "target__49870__176322",
                    "actions": {
                        "replace": [{
                            "source": "/ichibatop/com/inc/home/20080930/beta/tz/liquid/grouplist/20210416/pitari_non_sec.html",
                            "target": "group_banner_pitari"
                        }]
                    },
                    "match": {
                        "cu_secscv": 0
                    }
                }, {
                    "patternName": "target__target_card_scv__else_scv",
                    "patternId": "target__49870__176323",
                    "actions": {
                        "replace": [{
                            "source": "/ichibatop/com/inc/home/20080930/beta/tz/liquid/grouplist/20210416/pitari_else.html",
                            "target": "group_banner_pitari"
                        }]
                    }
                }],
                "type": "Target"
            },
            "match": {
                "cu_cardscv": 1
            }
        }, {
            "patternName": "target__root__non_card",
            "patternId": "target__49869__176325",
            "actions": {
                "replace": [{
                    "source": "/ichibatop/com/inc/home/20080930/beta/tz/liquid/grouplist/20210416/pitari_non_card.html",
                    "target": "group_banner_pitari"
                }]
            }
        }],
        "type": "Target"
    },
    "enableIframe": true,
    "experimentId": 9595,
    "cookieName": "Rp",
    "ratConfig": {
        "acc": 486,
        "aid": 1,
        "endPoint": "//rat.rakuten.co.jp/?cpkg_none=",
        "ckp": "Rz",
        "cks": "Rp"
    },
    "scvParams": {
        "acc": "1",
        "aid": "43",
        "attr": "111"
    },
    "campaignName": "ichiba_top_recommend_pc",
    "webViewEnabled": true,
    "apiType": "CCMP",
    "apiUrl": "https://user-attributes.api.rakuten.co.jp/capi/v1/ecosys/user.json"
});
(function() {
    if (document.addEventListener) {
        document.addEventListener('ph-CampaignEvent', function(event) {
            if (event.detail && event.detail.data) {
                var data = event.detail.data;
                data['phxpatternname'] = gP(data.phxpattern);
                if (9595 === data.phxexperiment) {
                    /**
                     * This function will be executed after your experiment processed
                     *
                     * Parameter references
                     * data.phxexperiment {string} Experiment ID
                     * data.phxpatternname {string} Matched pattern name
                     * data.phxsegments {array} Matched segment names
                     * data.phxabtestvalue {number} AB Test value
                     */

                }
            }
        });
    }

    function gP(p) {
        var t = '__';
        var s = p.split(t);
        return s[s.length - 1];
    }
})();