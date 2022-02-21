var __px = window.__px || {};
__px.campaigns = __px.campaigns || [];
__px.campaigns.push({
    "nest": {
        "conditions": [{
            "ratio": 95.0,
            "nest": {
                "conditions": [{
                    "nest": {
                        "conditions": [{
                            "ratio": 9.5,
                            "nest": {
                                "conditions": [{
                                    "ratio": 9.5,
                                    "patternName": "abtest__abtest_10__220127_career",
                                    "patternId": "abtest__52663__184342",
                                    "actions": {
                                        "replace": [{
                                            "source": "/ichibatop/com/inc/home/20080930/spt/inc/mno/pitari_mno_banner_01.tmpl?v=220127",
                                            "target": "mno-header"
                                        }]
                                    }
                                }],
                                "type": "ABTest"
                            }
                        }, {
                            "ratio": 95.0,
                            "nest": {
                                "conditions": [{
                                    "nest": {
                                        "conditions": [{
                                            "nest": {
                                                "conditions": [{
                                                    "nest": {
                                                        "conditions": [{
                                                            "patternName": "target__target_pay_user__220202_sec",
                                                            "patternId": "target__52667__184344",
                                                            "actions": {
                                                                "replace": [{
                                                                    "source": "/ichibatop/com/inc/home/20080930/spt/inc/mno/pitari_mno_banner_12.tmpl?v=220202",
                                                                    "target": "mno-header"
                                                                }]
                                                            },
                                                            "match": {
                                                                "up_rakuma": 1
                                                            }
                                                        }, {
                                                            "patternName": "target__target_pay_user__220214_rakuma",
                                                            "patternId": "target__52667__184345",
                                                            "actions": {
                                                                "replace": [{
                                                                    "source": "/ichibatop/com/inc/home/20080930/spt/inc/mno/pitari_mno_banner_03.tmpl?v=220214",
                                                                    "target": "mno-header"
                                                                }]
                                                            }
                                                        }],
                                                        "type": "Target"
                                                    },
                                                    "match": {
                                                        "up_pay": 1
                                                    }
                                                }, {
                                                    "patternName": "target__target_cardholder__220202_pay",
                                                    "patternId": "target__52666__184347",
                                                    "actions": {
                                                        "replace": [{
                                                            "source": "/ichibatop/com/inc/home/20080930/spt/inc/mno/pitari_mno_banner_11.tmpl?v=220202",
                                                            "target": "mno-header"
                                                        }]
                                                    }
                                                }],
                                                "type": "Target"
                                            },
                                            "match": {
                                                "up_card": 0
                                            }
                                        }, {
                                            "patternName": "target__target_bank_holder__220202_card",
                                            "patternId": "target__52665__184349",
                                            "actions": {
                                                "replace": [{
                                                    "source": "/ichibatop/com/inc/home/20080930/spt/inc/mno/pitari_mno_banner_10.tmpl?v=220202",
                                                    "target": "mno-header"
                                                }]
                                            }
                                        }],
                                        "type": "Target"
                                    },
                                    "match": {
                                        "cu_mbnk": 1
                                    }
                                }, {
                                    "patternName": "target__abtest_90__220202_bank",
                                    "patternId": "target__52664__184351",
                                    "actions": {
                                        "replace": [{
                                            "source": "/ichibatop/com/inc/home/20080930/spt/inc/mno/pitari_mno_banner_09.tmpl?v=220202",
                                            "target": "mno-header"
                                        }]
                                    }
                                }],
                                "type": "Target"
                            }
                        }],
                        "type": "ABTest"
                    },
                    "match": {
                        "co_mnoc": 1
                    }
                }, {
                    "nest": {
                        "conditions": [{
                            "ratio": 9.5,
                            "patternName": "abtest__target_else_segment__220204_bigs",
                            "patternId": "abtest__52668__184354",
                            "actions": {
                                "replace": [{
                                    "source": "/ichibatop/com/inc/home/20080930/spt/inc/mno/pitari_mno_banner_04.tmpl?v=220204",
                                    "target": "mno-header"
                                }]
                            }
                        }, {
                            "ratio": 19.0,
                            "patternName": "abtest__target_else_segment__220210_hand",
                            "patternId": "abtest__52668__184355",
                            "actions": {
                                "replace": [{
                                    "source": "/ichibatop/com/inc/home/20080930/spt/inc/mno/pitari_mno_banner_05.tmpl?v=220210",
                                    "target": "mno-header"
                                }]
                            }
                        }, {
                            "ratio": 95.0,
                            "nest": {
                                "conditions": [{
                                    "patternName": "target__abtest_80__220218_unlimit_ios",
                                    "patternId": "target__52669__184356",
                                    "actions": {
                                        "replace": [{
                                            "source": "/ichibatop/com/inc/home/20080930/spt/inc/mno/pitari_mno_banner_08.tmpl?v=211222",
                                            "target": "mno-header"
                                        }]
                                    },
                                    "match": {
                                        "cu_ios": 1
                                    }
                                }, {
                                    "nest": {
                                        "conditions": [{
                                            "ratio": 49.4,
                                            "patternName": "abtest__target_else_segment__211220_unlimit",
                                            "patternId": "abtest__52670__184357",
                                            "actions": {
                                                "replace": [{
                                                    "source": "/ichibatop/com/inc/home/20080930/spt/inc/mno/pitari_mno_banner_06.tmpl?v=211220",
                                                    "target": "mno-header"
                                                }]
                                            }
                                        }, {
                                            "ratio": 95.0,
                                            "patternName": "abtest__target_else_segment__211220_25000",
                                            "patternId": "abtest__52670__184358",
                                            "actions": {
                                                "replace": [{
                                                    "source": "/ichibatop/com/inc/home/20080930/spt/inc/mno/pitari_mno_banner_07.tmpl?v=211220",
                                                    "target": "mno-header"
                                                }]
                                            }
                                        }],
                                        "type": "ABTest"
                                    }
                                }],
                                "type": "Target"
                            }
                        }],
                        "type": "ABTest"
                    }
                }],
                "type": "Target"
            }
        }, {
            "ratio": 100.0,
            "patternName": "abtest__root__220202_unlimit_control",
            "patternId": "abtest__52660__184363",
            "actions": {
                "replace": [{
                    "source": "/ichibatop/com/inc/home/20080930/spt/inc/mno/pitari_mno_banner_13.tmpl?v=220202",
                    "target": "mno-header"
                }]
            }
        }],
        "type": "ABTest"
    },
    "enableIframe": true,
    "experimentId": 10105,
    "cookieName": "Rp",
    "ratConfig": {
        "acc": 1312,
        "aid": 1,
        "endPoint": "//secure.rat.rakuten.co.jp/?cpkg_none=",
        "ckp": "Rz",
        "cks": "Rp"
    },
    "scvParams": {
        "acc": "1",
        "aid": "43",
        "attr": "111"
    },
    "campaignName": "mnoprj_ichiba_top_pc",
    "webViewEnabled": true,
    "apiType": "CCMP",
    "apiUrl": "https://user-attributes.api.rakuten.co.jp/capi/v1/rmobile/user.json"
});