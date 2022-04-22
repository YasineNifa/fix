package dict

import (
	"github.com/quickfixgo/enum"
)

var PartyIDSources = map[enum.PartyIDSource]string{
	enum.PartyIDSource_KOREAN_INVESTOR_ID:                               "KOREAN_INVESTOR_ID",
	enum.PartyIDSource_TAIWANESE_QUALIFIED_FOREIGN_INVESTOR_ID_QFII_FID: "TAIWANESE_QUALIFIED_FOREIGN_INVESTOR_ID_QFII_FID",
	enum.PartyIDSource_TAIWANESE_TRADING_ACCT:                           "TAIWANESE_TRADING_ACCT",
	enum.PartyIDSource_MALAYSIAN_CENTRAL_DEPOSITORY:                     "MALAYSIAN_CENTRAL_DEPOSITORY",
	enum.PartyIDSource_CHINESE_INVESTOR_ID:                              "CHINESE_INVESTOR_ID",
	enum.PartyIDSource_UK_NATIONAL_INSURANCE_OR_PENSION_NUMBER:          "UK_NATIONAL_INSURANCE_OR_PENSION_NUMBER",
	enum.PartyIDSource_US_SOCIAL_SECURITY_NUMBER:                        "US_SOCIAL_SECURITY_NUMBER",
	enum.PartyIDSource_US_EMPLOYER_OR_TAX_ID_NUMBER:                     "US_EMPLOYER_OR_TAX_ID_NUMBER",
	enum.PartyIDSource_AUSTRALIAN_BUSINESS_NUMBER:                       "AUSTRALIAN_BUSINESS_NUMBER",
	enum.PartyIDSource_AUSTRALIAN_TAX_FILE_NUMBER:                       "AUSTRALIAN_TAX_FILE_NUMBER",
	enum.PartyIDSource_BIC:                                              "BIC",
	enum.PartyIDSource_GENERALLY_ACCEPTED_MARKET_PARTICIPANT_IDENTIFIER: "GENERALLY_ACCEPTED_MARKET_PARTICIPANT_IDENTIFIER",
	enum.PartyIDSource_PROPRIETARY:                                      "PROPRIETARY",
	enum.PartyIDSource_ISO_COUNTRY_CODE:                                 "ISO_COUNTRY_CODE",
	enum.PartyIDSource_SETTLEMENT_ENTITY_LOCATION:                       "SETTLEMENT_ENTITY_LOCATION",
	enum.PartyIDSource_MARKET_IDENTIFIER_CODE:                           "MARKET_IDENTIFIER_CODE",
	enum.PartyIDSource_CSD_PARTICIPANT_MEMBER_CODE:                      "CSD_PARTICIPANT_MEMBER_CODE",
	enum.PartyIDSource_DIRECTED_BROKER_THREE_CHARACTER_ACRONYM_AS_DEFINED_IN_ISITC_ETC_BEST_PRACTICE_GUIDELINES_DOCUMENT: "DIRECTED_BROKER_THREE_CHARACTER_ACRONYM_AS_DEFINED_IN_ISITC_ETC_BEST_PRACTICE_GUIDELINES_DOCUMENT",
}

var PartyIDSourcesReversed = map[string]enum.PartyIDSource{
	"KOREAN_INVESTOR_ID": enum.PartyIDSource_KOREAN_INVESTOR_ID,
	"TAIWANESE_QUALIFIED_FOREIGN_INVESTOR_ID_QFII_FID": enum.PartyIDSource_TAIWANESE_QUALIFIED_FOREIGN_INVESTOR_ID_QFII_FID,
	"TAIWANESE_TRADING_ACCT":                           enum.PartyIDSource_TAIWANESE_TRADING_ACCT,
	"MALAYSIAN_CENTRAL_DEPOSITORY":                     enum.PartyIDSource_MALAYSIAN_CENTRAL_DEPOSITORY,
	"CHINESE_INVESTOR_ID":                              enum.PartyIDSource_CHINESE_INVESTOR_ID,
	"UK_NATIONAL_INSURANCE_OR_PENSION_NUMBER":          enum.PartyIDSource_UK_NATIONAL_INSURANCE_OR_PENSION_NUMBER,
	"US_SOCIAL_SECURITY_NUMBER":                        enum.PartyIDSource_US_SOCIAL_SECURITY_NUMBER,
	"US_EMPLOYER_OR_TAX_ID_NUMBER":                     enum.PartyIDSource_US_EMPLOYER_OR_TAX_ID_NUMBER,
	"AUSTRALIAN_BUSINESS_NUMBER":                       enum.PartyIDSource_AUSTRALIAN_BUSINESS_NUMBER,
	"AUSTRALIAN_TAX_FILE_NUMBER":                       enum.PartyIDSource_AUSTRALIAN_TAX_FILE_NUMBER,
	"BIC":                                              enum.PartyIDSource_BIC,
	"GENERALLY_ACCEPTED_MARKET_PARTICIPANT_IDENTIFIER": enum.PartyIDSource_GENERALLY_ACCEPTED_MARKET_PARTICIPANT_IDENTIFIER,
	"PROPRIETARY":                                      enum.PartyIDSource_PROPRIETARY,
	"ISO_COUNTRY_CODE":                                 enum.PartyIDSource_ISO_COUNTRY_CODE,
	"SETTLEMENT_ENTITY_LOCATION":                       enum.PartyIDSource_SETTLEMENT_ENTITY_LOCATION,
	"MARKET_IDENTIFIER_CODE":                           enum.PartyIDSource_MARKET_IDENTIFIER_CODE,
	"CSD_PARTICIPANT_MEMBER_CODE":                      enum.PartyIDSource_CSD_PARTICIPANT_MEMBER_CODE,
	"DIRECTED_BROKER_THREE_CHARACTER_ACRONYM_AS_DEFINED_IN_ISITC_ETC_BEST_PRACTICE_GUIDELINES_DOCUMENT": enum.PartyIDSource_DIRECTED_BROKER_THREE_CHARACTER_ACRONYM_AS_DEFINED_IN_ISITC_ETC_BEST_PRACTICE_GUIDELINES_DOCUMENT,
}

var PartyRoles = map[enum.PartyRole]string{
	enum.PartyRole_EXECUTING_FIRM:                      "EXECUTING_FIRM",
	enum.PartyRole_SETTLEMENT_LOCATION:                 "SETTLEMENT_LOCATION",
	enum.PartyRole_ORDER_ORIGINATION_TRADER:            "ORDER_ORIGINATION_TRADER",
	enum.PartyRole_EXECUTING_TRADER:                    "EXECUTING_TRADER",
	enum.PartyRole_ORDER_ORIGINATION_FIRM:              "ORDER_ORIGINATION_FIRM",
	enum.PartyRole_GIVEUP_CLEARING_FIRM:                "GIVEUP_CLEARING_FIRM",
	enum.PartyRole_CORRESPONDANT_CLEARING_FIRM:         "CORRESPONDANT_CLEARING_FIRM",
	enum.PartyRole_EXECUTING_SYSTEM:                    "EXECUTING_SYSTEM",
	enum.PartyRole_CONTRA_FIRM:                         "CONTRA_FIRM",
	enum.PartyRole_CONTRA_CLEARING_FIRM:                "CONTRA_CLEARING_FIRM",
	enum.PartyRole_SPONSORING_FIRM:                     "SPONSORING_FIRM",
	enum.PartyRole_BROKER_OF_CREDIT:                    "BROKER_OF_CREDIT",
	enum.PartyRole_UNDERLYING_CONTRA_FIRM:              "UNDERLYING_CONTRA_FIRM",
	enum.PartyRole_CLEARING_ORGANIZATION:               "CLEARING_ORGANIZATION",
	enum.PartyRole_EXCHANGE:                            "EXCHANGE",
	enum.PartyRole_CUSTOMER_ACCOUNT:                    "CUSTOMER_ACCOUNT",
	enum.PartyRole_CORRESPONDENT_CLEARING_ORGANIZATION: "CORRESPONDENT_CLEARING_ORGANIZATION",
	enum.PartyRole_CORRESPONDENT_BROKER:                "CORRESPONDENT_BROKER",
	enum.PartyRole_BUYER_SELLER:                        "BUYER_SELLER",
	enum.PartyRole_CUSTODIAN:                           "CUSTODIAN",
	enum.PartyRole_INTERMEDIARY:                        "INTERMEDIARY",
	enum.PartyRole_CLIENT_ID:                           "CLIENT_ID",
	enum.PartyRole_AGENT:                               "AGENT",
	enum.PartyRole_SUB_CUSTODIAN:                       "SUB_CUSTODIAN",
	enum.PartyRole_BENEFICIARY:                         "BENEFICIARY",
	enum.PartyRole_INTERESTED_PARTY:                    "INTERESTED_PARTY",
	enum.PartyRole_REGULATORY_BODY:                     "REGULATORY_BODY",
	enum.PartyRole_LIQUIDITY_PROVIDER:                  "LIQUIDITY_PROVIDER",
	enum.PartyRole_ENTERING_TRADER:                     "ENTERING_TRADER",
	enum.PartyRole_CONTRA_TRADER:                       "CONTRA_TRADER",
	enum.PartyRole_POSITION_ACCOUNT:                    "POSITION_ACCOUNT",
	enum.PartyRole_CONTRA_INVESTOR_ID:                  "CONTRA_INVESTOR_ID",
	enum.PartyRole_CLEARING_FIRM:                       "CLEARING_FIRM",
	enum.PartyRole_TRANSFER_TO_FIRM:                    "TRANSFER_TO_FIRM",
	enum.PartyRole_CONTRA_POSITION_ACCOUNT:             "CONTRA_POSITION_ACCOUNT",
	enum.PartyRole_CONTRA_EXCHANGE:                     "CONTRA_EXCHANGE",
	enum.PartyRole_INTERNAL_CARRY_ACCOUNT:              "INTERNAL_CARRY_ACCOUNT",
	enum.PartyRole_ORDER_ENTRY_OPERATOR_ID:             "ORDER_ENTRY_OPERATOR_ID",
	enum.PartyRole_SECONDARY_ACCOUNT_NUMBER:            "SECONDARY_ACCOUNT_NUMBER",
	enum.PartyRole_FOREIGN_FIRM:                        "FOREIGN_FIRM",
	enum.PartyRole_THIRD_PARTY_ALLOCATION_FIRM:         "THIRD_PARTY_ALLOCATION_FIRM",
	enum.PartyRole_CLAIMING_ACCOUNT:                    "CLAIMING_ACCOUNT",
	enum.PartyRole_ASSET_MANAGER:                       "ASSET_MANAGER",
	enum.PartyRole_INVESTOR_ID:                         "INVESTOR_ID",
	enum.PartyRole_PLEDGOR_ACCOUNT:                     "PLEDGOR_ACCOUNT",
	enum.PartyRole_PLEDGEE_ACCOUNT:                     "PLEDGEE_ACCOUNT",
	enum.PartyRole_LARGE_TRADER_REPORTABLE_ACCOUNT:     "LARGE_TRADER_REPORTABLE_ACCOUNT",
	enum.PartyRole_TRADER_MNEMONIC:                     "TRADER_MNEMONIC",
	enum.PartyRole_SENDER_LOCATION:                     "SENDER_LOCATION",
	enum.PartyRole_SESSION_ID:                          "SESSION_ID",
	enum.PartyRole_ACCEPTABLE_COUNTERPARTY:             "ACCEPTABLE_COUNTERPARTY",
	enum.PartyRole_UNACCEPTABLE_COUNTERPARTY:           "UNACCEPTABLE_COUNTERPARTY",
	enum.PartyRole_ENTERING_UNIT:                       "ENTERING_UNIT",
	enum.PartyRole_EXECUTING_UNIT:                      "EXECUTING_UNIT",
	enum.PartyRole_INTRODUCING_FIRM:                    "INTRODUCING_FIRM",
	enum.PartyRole_INTRODUCING_BROKER:                  "INTRODUCING_BROKER",
	enum.PartyRole_QUOTE_ORIGINATOR:                    "QUOTE_ORIGINATOR",
	enum.PartyRole_REPORT_ORIGINATOR:                   "REPORT_ORIGINATOR",
	enum.PartyRole_SYSTEMATIC_INTERNALISER:             "SYSTEMATIC_INTERNALISER",
	enum.PartyRole_MULTILATERAL_TRADING_FACILITY:       "MULTILATERAL_TRADING_FACILITY",
	enum.PartyRole_REGULATED_MARKET:                    "REGULATED_MARKET",
	enum.PartyRole_MARKET_MAKER:                        "MARKET_MAKER",
	enum.PartyRole_INVESTMENT_FIRM:                     "INVESTMENT_FIRM",
	enum.PartyRole_HOST_COMPETENT_AUTHORITY:            "HOST_COMPETENT_AUTHORITY",
	enum.PartyRole_HOME_COMPETENT_AUTHORITY:            "HOME_COMPETENT_AUTHORITY",
	enum.PartyRole_ENTERING_FIRM:                       "ENTERING_FIRM",
	enum.PartyRole_COMPETENT_AUTHORITY_OF_THE_MOST_RELEVANT_MARKET_IN_TERMS_OF_LIQUIDITY: "COMPETENT_AUTHORITY_OF_THE_MOST_RELEVANT_MARKET_IN_TERMS_OF_LIQUIDITY",
	enum.PartyRole_COMPETENT_AUTHORITY_OF_THE_TRANSACTION:                                "COMPETENT_AUTHORITY_OF_THE_TRANSACTION",
	enum.PartyRole_REPORTING_INTERMEDIARY:                                                "REPORTING_INTERMEDIARY",
	enum.PartyRole_EXECUTION_VENUE:                                                       "EXECUTION_VENUE",
	enum.PartyRole_MARKET_DATA_ENTRY_ORIGINATOR:                                          "MARKET_DATA_ENTRY_ORIGINATOR",
	enum.PartyRole_LOCATION_ID:                                                           "LOCATION_ID",
	enum.PartyRole_DESK_ID:                                                               "DESK_ID",
	enum.PartyRole_MARKET_DATA_MARKET:                                                    "MARKET_DATA_MARKET",
	enum.PartyRole_ALLOCATION_ENTITY:                                                     "ALLOCATION_ENTITY",
	enum.PartyRole_PRIME_BROKER_PROVIDING_GENERAL_TRADE_SERVICES:                         "PRIME_BROKER_PROVIDING_GENERAL_TRADE_SERVICES",
	enum.PartyRole_LOCATE:                                                                "LOCATE",
	enum.PartyRole_STEP_OUT_FIRM:                                                         "STEP_OUT_FIRM",
	enum.PartyRole_BROKER_CIENT_ID:                                                       "BROKER_CIENT_ID",
	enum.PartyRole_CENTRAL_REGISTRATION_DEPOSITORY:                                       "CENTRAL_REGISTRATION_DEPOSITORY",
	enum.PartyRole_CLEARING_ACCOUNT:                                                      "CLEARING_ACCOUNT",
	enum.PartyRole_ACCEPTABLE_SETTLING_COUNTERPARTY:                                      "ACCEPTABLE_SETTLING_COUNTERPARTY",
	enum.PartyRole_UNACCEPTABLE_SETTLING_COUNTERPARTY:                                    "UNACCEPTABLE_SETTLING_COUNTERPARTY",
	enum.PartyRole_FUND_MANAGER_CLIENT_ID:                                                "FUND_MANAGER_CLIENT_ID",
}

var PartyRolesReversed = map[string]enum.PartyRole{
	"EXECUTING_FIRM":                      enum.PartyRole_EXECUTING_FIRM,
	"SETTLEMENT_LOCATION":                 enum.PartyRole_SETTLEMENT_LOCATION,
	"ORDER_ORIGINATION_TRADER":            enum.PartyRole_ORDER_ORIGINATION_TRADER,
	"EXECUTING_TRADER":                    enum.PartyRole_EXECUTING_TRADER,
	"ORDER_ORIGINATION_FIRM":              enum.PartyRole_ORDER_ORIGINATION_FIRM,
	"GIVEUP_CLEARING_FIRM":                enum.PartyRole_GIVEUP_CLEARING_FIRM,
	"CORRESPONDANT_CLEARING_FIRM":         enum.PartyRole_CORRESPONDANT_CLEARING_FIRM,
	"EXECUTING_SYSTEM":                    enum.PartyRole_EXECUTING_SYSTEM,
	"CONTRA_FIRM":                         enum.PartyRole_CONTRA_FIRM,
	"CONTRA_CLEARING_FIRM":                enum.PartyRole_CONTRA_CLEARING_FIRM,
	"SPONSORING_FIRM":                     enum.PartyRole_SPONSORING_FIRM,
	"BROKER_OF_CREDIT":                    enum.PartyRole_BROKER_OF_CREDIT,
	"UNDERLYING_CONTRA_FIRM":              enum.PartyRole_UNDERLYING_CONTRA_FIRM,
	"CLEARING_ORGANIZATION":               enum.PartyRole_CLEARING_ORGANIZATION,
	"EXCHANGE":                            enum.PartyRole_EXCHANGE,
	"CUSTOMER_ACCOUNT":                    enum.PartyRole_CUSTOMER_ACCOUNT,
	"CORRESPONDENT_CLEARING_ORGANIZATION": enum.PartyRole_CORRESPONDENT_CLEARING_ORGANIZATION,
	"CORRESPONDENT_BROKER":                enum.PartyRole_CORRESPONDENT_BROKER,
	"BUYER_SELLER":                        enum.PartyRole_BUYER_SELLER,
	"CUSTODIAN":                           enum.PartyRole_CUSTODIAN,
	"INTERMEDIARY":                        enum.PartyRole_INTERMEDIARY,
	"CLIENT_ID":                           enum.PartyRole_CLIENT_ID,
	"AGENT":                               enum.PartyRole_AGENT,
	"SUB_CUSTODIAN":                       enum.PartyRole_SUB_CUSTODIAN,
	"BENEFICIARY":                         enum.PartyRole_BENEFICIARY,
	"INTERESTED_PARTY":                    enum.PartyRole_INTERESTED_PARTY,
	"REGULATORY_BODY":                     enum.PartyRole_REGULATORY_BODY,
	"LIQUIDITY_PROVIDER":                  enum.PartyRole_LIQUIDITY_PROVIDER,
	"ENTERING_TRADER":                     enum.PartyRole_ENTERING_TRADER,
	"CONTRA_TRADER":                       enum.PartyRole_CONTRA_TRADER,
	"POSITION_ACCOUNT":                    enum.PartyRole_POSITION_ACCOUNT,
	"CONTRA_INVESTOR_ID":                  enum.PartyRole_CONTRA_INVESTOR_ID,
	"CLEARING_FIRM":                       enum.PartyRole_CLEARING_FIRM,
	"TRANSFER_TO_FIRM":                    enum.PartyRole_TRANSFER_TO_FIRM,
	"CONTRA_POSITION_ACCOUNT":             enum.PartyRole_CONTRA_POSITION_ACCOUNT,
	"CONTRA_EXCHANGE":                     enum.PartyRole_CONTRA_EXCHANGE,
	"INTERNAL_CARRY_ACCOUNT":              enum.PartyRole_INTERNAL_CARRY_ACCOUNT,
	"ORDER_ENTRY_OPERATOR_ID":             enum.PartyRole_ORDER_ENTRY_OPERATOR_ID,
	"SECONDARY_ACCOUNT_NUMBER":            enum.PartyRole_SECONDARY_ACCOUNT_NUMBER,
	"FOREIGN_FIRM":                        enum.PartyRole_FOREIGN_FIRM,
	"THIRD_PARTY_ALLOCATION_FIRM":         enum.PartyRole_THIRD_PARTY_ALLOCATION_FIRM,
	"CLAIMING_ACCOUNT":                    enum.PartyRole_CLAIMING_ACCOUNT,
	"ASSET_MANAGER":                       enum.PartyRole_ASSET_MANAGER,
	"INVESTOR_ID":                         enum.PartyRole_INVESTOR_ID,
	"PLEDGOR_ACCOUNT":                     enum.PartyRole_PLEDGOR_ACCOUNT,
	"PLEDGEE_ACCOUNT":                     enum.PartyRole_PLEDGEE_ACCOUNT,
	"LARGE_TRADER_REPORTABLE_ACCOUNT":     enum.PartyRole_LARGE_TRADER_REPORTABLE_ACCOUNT,
	"TRADER_MNEMONIC":                     enum.PartyRole_TRADER_MNEMONIC,
	"SENDER_LOCATION":                     enum.PartyRole_SENDER_LOCATION,
	"SESSION_ID":                          enum.PartyRole_SESSION_ID,
	"ACCEPTABLE_COUNTERPARTY":             enum.PartyRole_ACCEPTABLE_COUNTERPARTY,
	"UNACCEPTABLE_COUNTERPARTY":           enum.PartyRole_UNACCEPTABLE_COUNTERPARTY,
	"ENTERING_UNIT":                       enum.PartyRole_ENTERING_UNIT,
	"EXECUTING_UNIT":                      enum.PartyRole_EXECUTING_UNIT,
	"INTRODUCING_FIRM":                    enum.PartyRole_INTRODUCING_FIRM,
	"INTRODUCING_BROKER":                  enum.PartyRole_INTRODUCING_BROKER,
	"QUOTE_ORIGINATOR":                    enum.PartyRole_QUOTE_ORIGINATOR,
	"REPORT_ORIGINATOR":                   enum.PartyRole_REPORT_ORIGINATOR,
	"SYSTEMATIC_INTERNALISER":             enum.PartyRole_SYSTEMATIC_INTERNALISER,
	"MULTILATERAL_TRADING_FACILITY":       enum.PartyRole_MULTILATERAL_TRADING_FACILITY,
	"REGULATED_MARKET":                    enum.PartyRole_REGULATED_MARKET,
	"MARKET_MAKER":                        enum.PartyRole_MARKET_MAKER,
	"INVESTMENT_FIRM":                     enum.PartyRole_INVESTMENT_FIRM,
	"HOST_COMPETENT_AUTHORITY":            enum.PartyRole_HOST_COMPETENT_AUTHORITY,
	"HOME_COMPETENT_AUTHORITY":            enum.PartyRole_HOME_COMPETENT_AUTHORITY,
	"ENTERING_FIRM":                       enum.PartyRole_ENTERING_FIRM,
	"COMPETENT_AUTHORITY_OF_THE_MOST_RELEVANT_MARKET_IN_TERMS_OF_LIQUIDITY": enum.PartyRole_COMPETENT_AUTHORITY_OF_THE_MOST_RELEVANT_MARKET_IN_TERMS_OF_LIQUIDITY,
	"COMPETENT_AUTHORITY_OF_THE_TRANSACTION":                                enum.PartyRole_COMPETENT_AUTHORITY_OF_THE_TRANSACTION,
	"REPORTING_INTERMEDIARY":                                                enum.PartyRole_REPORTING_INTERMEDIARY,
	"EXECUTION_VENUE":                                                       enum.PartyRole_EXECUTION_VENUE,
	"MARKET_DATA_ENTRY_ORIGINATOR":                                          enum.PartyRole_MARKET_DATA_ENTRY_ORIGINATOR,
	"LOCATION_ID":                                                           enum.PartyRole_LOCATION_ID,
	"DESK_ID":                                                               enum.PartyRole_DESK_ID,
	"MARKET_DATA_MARKET":                                                    enum.PartyRole_MARKET_DATA_MARKET,
	"ALLOCATION_ENTITY":                                                     enum.PartyRole_ALLOCATION_ENTITY,
	"PRIME_BROKER_PROVIDING_GENERAL_TRADE_SERVICES":                         enum.PartyRole_PRIME_BROKER_PROVIDING_GENERAL_TRADE_SERVICES,
	"LOCATE":                             enum.PartyRole_LOCATE,
	"STEP_OUT_FIRM":                      enum.PartyRole_STEP_OUT_FIRM,
	"BROKER_CIENT_ID":                    enum.PartyRole_BROKER_CIENT_ID,
	"CENTRAL_REGISTRATION_DEPOSITORY":    enum.PartyRole_CENTRAL_REGISTRATION_DEPOSITORY,
	"CLEARING_ACCOUNT":                   enum.PartyRole_CLEARING_ACCOUNT,
	"ACCEPTABLE_SETTLING_COUNTERPARTY":   enum.PartyRole_ACCEPTABLE_SETTLING_COUNTERPARTY,
	"UNACCEPTABLE_SETTLING_COUNTERPARTY": enum.PartyRole_UNACCEPTABLE_SETTLING_COUNTERPARTY,
	"FUND_MANAGER_CLIENT_ID":             enum.PartyRole_FUND_MANAGER_CLIENT_ID,
}

var PartySubIDTypes = map[enum.PartySubIDType]string{
	enum.PartySubIDType_FIRM:                                                                      "FIRM",
	enum.PartySubIDType_SECURITIES_ACCOUNT_NUMBER:                                                 "SECURITIES_ACCOUNT_NUMBER",
	enum.PartySubIDType_REGISTRATION_NUMBER:                                                       "REGISTRATION_NUMBER",
	enum.PartySubIDType_REGISTERED_ADDRESS_12:                                                     "REGISTERED_ADDRESS_12",
	enum.PartySubIDType_REGULATORY_STATUS:                                                         "REGULATORY_STATUS",
	enum.PartySubIDType_REGISTRATION_NAME:                                                         "REGISTRATION_NAME",
	enum.PartySubIDType_CASH_ACCOUNT_NUMBER:                                                       "CASH_ACCOUNT_NUMBER",
	enum.PartySubIDType_BIC:                                                                       "BIC",
	enum.PartySubIDType_CSD_PARTICIPANT_MEMBER_CODE:                                               "CSD_PARTICIPANT_MEMBER_CODE",
	enum.PartySubIDType_REGISTERED_ADDRESS_18:                                                     "REGISTERED_ADDRESS_18",
	enum.PartySubIDType_FUND_ACCOUNT_NAME:                                                         "FUND_ACCOUNT_NAME",
	enum.PartySubIDType_PERSON:                                                                    "PERSON",
	enum.PartySubIDType_TELEX_NUMBER:                                                              "TELEX_NUMBER",
	enum.PartySubIDType_FAX_NUMBER:                                                                "FAX_NUMBER",
	enum.PartySubIDType_SECURITIES_ACCOUNT_NAME:                                                   "SECURITIES_ACCOUNT_NAME",
	enum.PartySubIDType_CASH_ACCOUNT_NAME:                                                         "CASH_ACCOUNT_NAME",
	enum.PartySubIDType_DEPARTMENT:                                                                "DEPARTMENT",
	enum.PartySubIDType_LOCATION_DESK:                                                             "LOCATION_DESK",
	enum.PartySubIDType_POSITION_ACCOUNT_TYPE:                                                     "POSITION_ACCOUNT_TYPE",
	enum.PartySubIDType_SECURITY_LOCATE_ID:                                                        "SECURITY_LOCATE_ID",
	enum.PartySubIDType_MARKET_MAKER:                                                              "MARKET_MAKER",
	enum.PartySubIDType_ELIGIBLE_COUNTERPARTY:                                                     "ELIGIBLE_COUNTERPARTY",
	enum.PartySubIDType_SYSTEM:                                                                    "SYSTEM",
	enum.PartySubIDType_PROFESSIONAL_CLIENT:                                                       "PROFESSIONAL_CLIENT",
	enum.PartySubIDType_LOCATION:                                                                  "LOCATION",
	enum.PartySubIDType_EXECUTION_VENUE:                                                           "EXECUTION_VENUE",
	enum.PartySubIDType_CURRENCY_DELIVERY_IDENTIFIER:                                              "CURRENCY_DELIVERY_IDENTIFIER",
	enum.PartySubIDType_ADDRESS_CITY:                                                              "ADDRESS_CITY",
	enum.PartySubIDType_ADDRESS_STATE_PROVINCE:                                                    "ADDRESS_STATE_PROVINCE",
	enum.PartySubIDType_ADDRESS_POSTAL_CODE:                                                       "ADDRESS_POSTAL_CODE",
	enum.PartySubIDType_ADDRESS_STREET:                                                            "ADDRESS_STREET",
	enum.PartySubIDType_ADDRESS_COUNTRY:                                                           "ADDRESS_COUNTRY",
	enum.PartySubIDType_ISO_COUNTRY_CODE:                                                          "ISO_COUNTRY_CODE",
	enum.PartySubIDType_APPLICATION:                                                               "APPLICATION",
	enum.PartySubIDType_MARKET_SEGMENT:                                                            "MARKET_SEGMENT",
	enum.PartySubIDType_RESERVEDANDAVAILABLEFORBILATERALLYAGREEDUPONUSERDEFINEDVALUES:             "RESERVEDANDAVAILABLEFORBILATERALLYAGREEDUPONUSERDEFINEDVALUES",
	enum.PartySubIDType_CUSTOMER_ACCOUNT_TYPE:                                                     "CUSTOMER_ACCOUNT_TYPE",
	enum.PartySubIDType_OMNIBUS_ACCOUNT:                                                           "OMNIBUS_ACCOUNT",
	enum.PartySubIDType_FUNDS_SEGREGATION_TYPE:                                                    "FUNDS_SEGREGATION_TYPE",
	enum.PartySubIDType_GUARANTEE_FUND:                                                            "GUARANTEE_FUND",
	enum.PartySubIDType_SWAP_DEALER:                                                               "SWAP_DEALER",
	enum.PartySubIDType_MAJOR_PARTICIPANT:                                                         "MAJOR_PARTICIPANT",
	enum.PartySubIDType_FINANCIAL_ENTITY:                                                          "FINANCIAL_ENTITY",
	enum.PartySubIDType_US_PERSON:                                                                 "US_PERSON",
	enum.PartySubIDType_REPORTING_ENTITY_INDICATOR:                                                "REPORTING_ENTITY_INDICATOR",
	enum.PartySubIDType_FULL_LEGAL_NAME_OF_FIRM:                                                   "FULL_LEGAL_NAME_OF_FIRM",
	enum.PartySubIDType_ELECTED_CLEARING_REQUIREMENT_EXCEPTION:                                    "ELECTED_CLEARING_REQUIREMENT_EXCEPTION",
	enum.PartySubIDType_BUSINESS_CENTER:                                                           "BUSINESS_CENTER",
	enum.PartySubIDType_REFERENCE_TEXT:                                                            "REFERENCE_TEXT",
	enum.PartySubIDType_SHORT_MARKING_EXEMPT_ACCOUNT:                                              "SHORT_MARKING_EXEMPT_ACCOUNT",
	enum.PartySubIDType_PARENT_FIRM_IDENTIFIER:                                                    "PARENT_FIRM_IDENTIFIER",
	enum.PartySubIDType_PARENT_FIRM_NAME:                                                          "PARENT_FIRM_NAME",
	enum.PartySubIDType_DEAL_IDENTIFIER:                                                           "DEAL_IDENTIFIER",
	enum.PartySubIDType_SYSTEM_TRADE_IDENTIFIER:                                                   "SYSTEM_TRADE_IDENTIFIER",
	enum.PartySubIDType_SYSTEM_TRADE_SUB_IDENTIFIER:                                               "SYSTEM_TRADE_SUB_IDENTIFIER",
	enum.PartySubIDType_FUTURES_COMMISSION_MERCHANT:                                               "FUTURES_COMMISSION_MERCHANT",
	enum.PartySubIDType_POSTAL_ADDRESS:                                                            "POSTAL_ADDRESS",
	enum.PartySubIDType_DELIVERY_TERMINAL_CUSTOMER_ACCOUNT_CODE:                                   "DELIVERY_TERMINAL_CUSTOMER_ACCOUNT_CODE",
	enum.PartySubIDType_VOLUNTARY_REPORTING_ENTITY:                                                "VOLUNTARY_REPORTING_ENTITY",
	enum.PartySubIDType_REPORTING_OBLIGATION_JURISDICTION:                                         "REPORTING_OBLIGATION_JURISDICTION",
	enum.PartySubIDType_VOLUNTARY_REPORTING_JURISDICTION:                                          "VOLUNTARY_REPORTING_JURISDICTION",
	enum.PartySubIDType_COMPANY_ACTIVITIES:                                                        "COMPANY_ACTIVITIES",
	enum.PartySubIDType_EUROPEAN_ECONOMIC_AREA_DOMICILED:                                          "EUROPEAN_ECONOMIC_AREA_DOMICILED",
	enum.PartySubIDType_CONTRACT_LINKED_TO_COMMERCIAL_OR_TREASURY_FINANCING_FOR_THIS_COUNTERPARTY: "CONTRACT_LINKED_TO_COMMERCIAL_OR_TREASURY_FINANCING_FOR_THIS_COUNTERPARTY",
	enum.PartySubIDType_CONTRACT_ABOVE_CLEARING_THRESHOLD_FOR_THIS_COUNTERPARTY:                   "CONTRACT_ABOVE_CLEARING_THRESHOLD_FOR_THIS_COUNTERPARTY",
	enum.PartySubIDType_VOLUNTARY_REPORTING_PARTY:                                                 "VOLUNTARY_REPORTING_PARTY",
	enum.PartySubIDType_END_USER:                                                                  "END_USER",
	enum.PartySubIDType_PHONE_NUMBER:                                                              "PHONE_NUMBER",
	enum.PartySubIDType_LOCATION_OR_JURISDICTION:                                                  "LOCATION_OR_JURISDICTION",
	enum.PartySubIDType_DERIVATIVES_DEALER:                                                        "DERIVATIVES_DEALER",
	enum.PartySubIDType_DOMICILE:                                                                  "DOMICILE",
	enum.PartySubIDType_EXEMPT_FROM_RECOGNITION:                                                   "EXEMPT_FROM_RECOGNITION",
	enum.PartySubIDType_PAYER:                                                                     "PAYER",
	enum.PartySubIDType_RECEIVER:                                                                  "RECEIVER",
	enum.PartySubIDType_SYSTEMATIC_INTERNALISER:                                                   "SYSTEMATIC_INTERNALISER",
	enum.PartySubIDType_PUBLISHING_ENTITY_INDICATOR:                                               "PUBLISHING_ENTITY_INDICATOR",
	enum.PartySubIDType_FIRST_NAME:                                                                "FIRST_NAME",
	enum.PartySubIDType_SURNAME:                                                                   "SURNAME",
	enum.PartySubIDType_EMAIL_ADDRESS:                                                             "EMAIL_ADDRESS",
	enum.PartySubIDType_DATE_OF_BIRTH:                                                             "DATE_OF_BIRTH",
	enum.PartySubIDType_ORDER_TRANSMITTING_FIRM:                                                   "ORDER_TRANSMITTING_FIRM",
	enum.PartySubIDType_ORDER_TRANSMITTING_FIRM_FOR_BUYER:                                         "ORDER_TRANSMITTING_FIRM_FOR_BUYER",
	enum.PartySubIDType_ORDER_TRANSMITTER_FOR_SELLER:                                              "ORDER_TRANSMITTER_FOR_SELLER",
	enum.PartySubIDType_LEGAL_ENTITY_IDENTIFIER:                                                   "LEGAL_ENTITY_IDENTIFIER",
	enum.PartySubIDType_SUB_SECTOR_CLASSIFICATION:                                                 "SUB_SECTOR_CLASSIFICATION",
	enum.PartySubIDType_PARTY_SIDE:                                                                "PARTY_SIDE",
	enum.PartySubIDType_LEGAL_REGISTRATION_COUNTRY:                                                "LEGAL_REGISTRATION_COUNTRY",
	enum.PartySubIDType_CONTACT_NAME:                                                              "CONTACT_NAME",
}

var PartySubIDTypesReversed = map[string]enum.PartySubIDType{
	"FIRM":                         enum.PartySubIDType_FIRM,
	"SECURITIES_ACCOUNT_NUMBER":    enum.PartySubIDType_SECURITIES_ACCOUNT_NUMBER,
	"REGISTRATION_NUMBER":          enum.PartySubIDType_REGISTRATION_NUMBER,
	"REGISTERED_ADDRESS_12":        enum.PartySubIDType_REGISTERED_ADDRESS_12,
	"REGULATORY_STATUS":            enum.PartySubIDType_REGULATORY_STATUS,
	"REGISTRATION_NAME":            enum.PartySubIDType_REGISTRATION_NAME,
	"CASH_ACCOUNT_NUMBER":          enum.PartySubIDType_CASH_ACCOUNT_NUMBER,
	"BIC":                          enum.PartySubIDType_BIC,
	"CSD_PARTICIPANT_MEMBER_CODE":  enum.PartySubIDType_CSD_PARTICIPANT_MEMBER_CODE,
	"REGISTERED_ADDRESS_18":        enum.PartySubIDType_REGISTERED_ADDRESS_18,
	"FUND_ACCOUNT_NAME":            enum.PartySubIDType_FUND_ACCOUNT_NAME,
	"PERSON":                       enum.PartySubIDType_PERSON,
	"TELEX_NUMBER":                 enum.PartySubIDType_TELEX_NUMBER,
	"FAX_NUMBER":                   enum.PartySubIDType_FAX_NUMBER,
	"SECURITIES_ACCOUNT_NAME":      enum.PartySubIDType_SECURITIES_ACCOUNT_NAME,
	"CASH_ACCOUNT_NAME":            enum.PartySubIDType_CASH_ACCOUNT_NAME,
	"DEPARTMENT":                   enum.PartySubIDType_DEPARTMENT,
	"LOCATION_DESK":                enum.PartySubIDType_LOCATION_DESK,
	"POSITION_ACCOUNT_TYPE":        enum.PartySubIDType_POSITION_ACCOUNT_TYPE,
	"SECURITY_LOCATE_ID":           enum.PartySubIDType_SECURITY_LOCATE_ID,
	"MARKET_MAKER":                 enum.PartySubIDType_MARKET_MAKER,
	"ELIGIBLE_COUNTERPARTY":        enum.PartySubIDType_ELIGIBLE_COUNTERPARTY,
	"SYSTEM":                       enum.PartySubIDType_SYSTEM,
	"PROFESSIONAL_CLIENT":          enum.PartySubIDType_PROFESSIONAL_CLIENT,
	"LOCATION":                     enum.PartySubIDType_LOCATION,
	"EXECUTION_VENUE":              enum.PartySubIDType_EXECUTION_VENUE,
	"CURRENCY_DELIVERY_IDENTIFIER": enum.PartySubIDType_CURRENCY_DELIVERY_IDENTIFIER,
	"ADDRESS_CITY":                 enum.PartySubIDType_ADDRESS_CITY,
	"ADDRESS_STATE_PROVINCE":       enum.PartySubIDType_ADDRESS_STATE_PROVINCE,
	"ADDRESS_POSTAL_CODE":          enum.PartySubIDType_ADDRESS_POSTAL_CODE,
	"ADDRESS_STREET":               enum.PartySubIDType_ADDRESS_STREET,
	"ADDRESS_COUNTRY":              enum.PartySubIDType_ADDRESS_COUNTRY,
	"ISO_COUNTRY_CODE":             enum.PartySubIDType_ISO_COUNTRY_CODE,
	"APPLICATION":                  enum.PartySubIDType_APPLICATION,
	"MARKET_SEGMENT":               enum.PartySubIDType_MARKET_SEGMENT,
	"RESERVEDANDAVAILABLEFORBILATERALLYAGREEDUPONUSERDEFINEDVALUES": enum.PartySubIDType_RESERVEDANDAVAILABLEFORBILATERALLYAGREEDUPONUSERDEFINEDVALUES,
	"CUSTOMER_ACCOUNT_TYPE":                   enum.PartySubIDType_CUSTOMER_ACCOUNT_TYPE,
	"OMNIBUS_ACCOUNT":                         enum.PartySubIDType_OMNIBUS_ACCOUNT,
	"FUNDS_SEGREGATION_TYPE":                  enum.PartySubIDType_FUNDS_SEGREGATION_TYPE,
	"GUARANTEE_FUND":                          enum.PartySubIDType_GUARANTEE_FUND,
	"SWAP_DEALER":                             enum.PartySubIDType_SWAP_DEALER,
	"MAJOR_PARTICIPANT":                       enum.PartySubIDType_MAJOR_PARTICIPANT,
	"FINANCIAL_ENTITY":                        enum.PartySubIDType_FINANCIAL_ENTITY,
	"US_PERSON":                               enum.PartySubIDType_US_PERSON,
	"REPORTING_ENTITY_INDICATOR":              enum.PartySubIDType_REPORTING_ENTITY_INDICATOR,
	"FULL_LEGAL_NAME_OF_FIRM":                 enum.PartySubIDType_FULL_LEGAL_NAME_OF_FIRM,
	"ELECTED_CLEARING_REQUIREMENT_EXCEPTION":  enum.PartySubIDType_ELECTED_CLEARING_REQUIREMENT_EXCEPTION,
	"BUSINESS_CENTER":                         enum.PartySubIDType_BUSINESS_CENTER,
	"REFERENCE_TEXT":                          enum.PartySubIDType_REFERENCE_TEXT,
	"SHORT_MARKING_EXEMPT_ACCOUNT":            enum.PartySubIDType_SHORT_MARKING_EXEMPT_ACCOUNT,
	"PARENT_FIRM_IDENTIFIER":                  enum.PartySubIDType_PARENT_FIRM_IDENTIFIER,
	"PARENT_FIRM_NAME":                        enum.PartySubIDType_PARENT_FIRM_NAME,
	"DEAL_IDENTIFIER":                         enum.PartySubIDType_DEAL_IDENTIFIER,
	"SYSTEM_TRADE_IDENTIFIER":                 enum.PartySubIDType_SYSTEM_TRADE_IDENTIFIER,
	"SYSTEM_TRADE_SUB_IDENTIFIER":             enum.PartySubIDType_SYSTEM_TRADE_SUB_IDENTIFIER,
	"FUTURES_COMMISSION_MERCHANT":             enum.PartySubIDType_FUTURES_COMMISSION_MERCHANT,
	"POSTAL_ADDRESS":                          enum.PartySubIDType_POSTAL_ADDRESS,
	"DELIVERY_TERMINAL_CUSTOMER_ACCOUNT_CODE": enum.PartySubIDType_DELIVERY_TERMINAL_CUSTOMER_ACCOUNT_CODE,
	"VOLUNTARY_REPORTING_ENTITY":              enum.PartySubIDType_VOLUNTARY_REPORTING_ENTITY,
	"REPORTING_OBLIGATION_JURISDICTION":       enum.PartySubIDType_REPORTING_OBLIGATION_JURISDICTION,
	"VOLUNTARY_REPORTING_JURISDICTION":        enum.PartySubIDType_VOLUNTARY_REPORTING_JURISDICTION,
	"COMPANY_ACTIVITIES":                      enum.PartySubIDType_COMPANY_ACTIVITIES,
	"EUROPEAN_ECONOMIC_AREA_DOMICILED":        enum.PartySubIDType_EUROPEAN_ECONOMIC_AREA_DOMICILED,
	"CONTRACT_LINKED_TO_COMMERCIAL_OR_TREASURY_FINANCING_FOR_THIS_COUNTERPARTY": enum.PartySubIDType_CONTRACT_LINKED_TO_COMMERCIAL_OR_TREASURY_FINANCING_FOR_THIS_COUNTERPARTY,
	"CONTRACT_ABOVE_CLEARING_THRESHOLD_FOR_THIS_COUNTERPARTY":                   enum.PartySubIDType_CONTRACT_ABOVE_CLEARING_THRESHOLD_FOR_THIS_COUNTERPARTY,
	"VOLUNTARY_REPORTING_PARTY":                                                 enum.PartySubIDType_VOLUNTARY_REPORTING_PARTY,
	"END_USER":                                                                  enum.PartySubIDType_END_USER,
	"PHONE_NUMBER":                                                              enum.PartySubIDType_PHONE_NUMBER,
	"LOCATION_OR_JURISDICTION":                                                  enum.PartySubIDType_LOCATION_OR_JURISDICTION,
	"DERIVATIVES_DEALER":                                                        enum.PartySubIDType_DERIVATIVES_DEALER,
	"DOMICILE":                                                                  enum.PartySubIDType_DOMICILE,
	"EXEMPT_FROM_RECOGNITION":                                                   enum.PartySubIDType_EXEMPT_FROM_RECOGNITION,
	"PAYER":                                                                     enum.PartySubIDType_PAYER,
	"RECEIVER":                                                                  enum.PartySubIDType_RECEIVER,
	"SYSTEMATIC_INTERNALISER":                                                   enum.PartySubIDType_SYSTEMATIC_INTERNALISER,
	"PUBLISHING_ENTITY_INDICATOR":                                               enum.PartySubIDType_PUBLISHING_ENTITY_INDICATOR,
	"FIRST_NAME":                                                                enum.PartySubIDType_FIRST_NAME,
	"SURNAME":                                                                   enum.PartySubIDType_SURNAME,
	"EMAIL_ADDRESS":                                                             enum.PartySubIDType_EMAIL_ADDRESS,
	"DATE_OF_BIRTH":                                                             enum.PartySubIDType_DATE_OF_BIRTH,
	"ORDER_TRANSMITTING_FIRM":                                                   enum.PartySubIDType_ORDER_TRANSMITTING_FIRM,
	"ORDER_TRANSMITTING_FIRM_FOR_BUYER":                                         enum.PartySubIDType_ORDER_TRANSMITTING_FIRM_FOR_BUYER,
	"ORDER_TRANSMITTER_FOR_SELLER":                                              enum.PartySubIDType_ORDER_TRANSMITTER_FOR_SELLER,
	"LEGAL_ENTITY_IDENTIFIER":                                                   enum.PartySubIDType_LEGAL_ENTITY_IDENTIFIER,
	"SUB_SECTOR_CLASSIFICATION":                                                 enum.PartySubIDType_SUB_SECTOR_CLASSIFICATION,
	"PARTY_SIDE":                                                                enum.PartySubIDType_PARTY_SIDE,
	"LEGAL_REGISTRATION_COUNTRY":                                                enum.PartySubIDType_LEGAL_REGISTRATION_COUNTRY,
	"CONTACT_NAME":                                                              enum.PartySubIDType_CONTACT_NAME,
}
