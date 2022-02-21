var __px = window.__px || {};
__px.campaigns = __px.campaigns || [];
__px.campaigns.push({
    "nest": {
        "conditions": [{
            "patternName": "target__root__rtc_regular_visa",
            "patternId": "target__49985__176601",
            "actions": {
                "replace": [{
                    "source": "/ichibatop/com/inc/home/20080930/beta/tz/opt/right_upper_banner/20201027_kobo/pitari_rtc_nonholder_regular.html",
                    "target": "right_upper_banner"
                }]
            },
            "match": {
                "__pitari_non_login": 0
            }
        }, {
            "nest": {
                "conditions": [{
                    "patternName": "target__target_card_scv__rmb_mno_2009",
                    "patternId": "target__49986__176602",
                    "actions": {
                        "replace": [{
                            "source": "/ichibatop/com/inc/home/20080930/beta/tz/opt/right_upper_banner/20200911_kobo/pitari_rmb_mno_2009.html",
                            "target": "right_upper_banner"
                        }]
                    },
                    "match": {
                        "cu_mnoscv": 0
                    }
                }, {
                    "patternName": "target__target_card_scv__bnk_1000",
                    "patternId": "target__49986__176603",
                    "actions": {
                        "replace": [{
                            "source": "/ichibatop/com/inc/home/20080930/beta/tz/opt/right_upper_banner/20201111_kobo/pitari_bank_1000.html",
                            "target": "right_upper_banner"
                        }]
                    },
                    "match": {
                        "cu_bankscv": 0
                    }
                }, {
                    "patternName": "target__target_card_scv__keiba_1000",
                    "patternId": "target__49986__176604",
                    "actions": {
                        "replace": [{
                            "source": "/ichibatop/com/inc/home/20080930/beta/tz/opt/right_upper_banner/20201111_kobo/pitari_keiba_1000.html",
                            "target": "right_upper_banner"
                        }]
                    },
                    "match": {
                        "cu_keibascv": 0
                    }
                }, {
                    "patternName": "target__target_card_scv__card_ex",
                    "patternId": "target__49986__176605",
                    "actions": {
                        "replace": [{
                            "source": "/ichibatop/com/inc/home/20080930/beta/tz/opt/right_upper_banner/20210712_kobo/pitari_card_ex.html",
                            "target": "right_upper_banner"
                        }]
                    }
                }],
                "type": "Target"
            },
            "match": {
                "cu_cardscv": 1
            }
        }, {
            "nest": {
                "conditions": [{
                    "patternName": "target__target_else__rmb_mno_2009",
                    "patternId": "target__49987__176607",
                    "actions": {
                        "replace": [{
                            "source": "/ichibatop/com/inc/home/20080930/beta/tz/opt/right_upper_banner/20200911_kobo/pitari_rmb_mno_2009.html",
                            "target": "right_upper_banner"
                        }]
                    },
                    "match": {
                        "co_user_rank": 7
                    }
                }, {
                    "patternName": "target__target_else__rtc_platinum_visa",
                    "patternId": "target__49987__176608",
                    "actions": {
                        "replace": [{
                            "source": "/ichibatop/com/inc/home/20080930/beta/tz/opt/right_upper_banner/20201027_kobo/pitari_rtc_nonholder_platinum.html",
                            "target": "right_upper_banner"
                        }]
                    },
                    "match": {
                        "co_user_rank": 6
                    }
                }, {
                    "patternName": "target__target_else__rtc_gold_visa",
                    "patternId": "target__49987__176609",
                    "actions": {
                        "replace": [{
                            "source": "/ichibatop/com/inc/home/20080930/beta/tz/opt/right_upper_banner/20201027_kobo/pitari_rtc_nonholder_gold.html",
                            "target": "right_upper_banner"
                        }]
                    },
                    "match": {
                        "co_user_rank": 5
                    }
                }, {
                    "patternName": "target__target_else__rtc_silver_visa",
                    "patternId": "target__49987__176610",
                    "actions": {
                        "replace": [{
                            "source": "/ichibatop/com/inc/home/20080930/beta/tz/opt/right_upper_banner/20201027_kobo/pitari_rtc_nonholder_silver.html",
                            "target": "right_upper_banner"
                        }]
                    },
                    "match": {
                        "co_user_rank": 4
                    }
                }, {
                    "patternName": "target__target_else__rtc_regular_visa",
                    "patternId": "target__49987__176611",
                    "actions": {
                        "replace": [{
                            "source": "/ichibatop/com/inc/home/20080930/beta/tz/opt/right_upper_banner/20201027_kobo/pitari_rtc_nonholder_regular.html",
                            "target": "right_upper_banner"
                        }]
                    }
                }],
                "type": "Target"
            }
        }],
        "type": "Target"
    },
    "enableIframe": true,
    "experimentId": 9615,
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
    "campaignName": "ichiba_top_page_right_banner_pc",
    "webViewEnabled": true,
    "apiType": "CCMP",
    "apiUrl": "https://user-attributes.api.rakuten.co.jp/capi/v1/ecosys/user.json"
});