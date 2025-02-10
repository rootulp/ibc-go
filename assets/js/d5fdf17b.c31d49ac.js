"use strict";(self.webpackChunkdocs=self.webpackChunkdocs||[]).push([[69834],{62801:(e,t,s)=>{s.r(t),s.d(t,{assets:()=>r,contentTitle:()=>i,default:()=>h,frontMatter:()=>a,metadata:()=>c,toc:()=>l});var n=s(85893),o=s(11151);const a={title:"Governance",sidebar_label:"Governance",sidebar_position:5,slug:"/ibc/light-clients/wasm/governance"},i="Governance",c={id:"light-clients/wasm/governance",title:"Governance",description:"Learn how to upload Wasm light client byte code on a chain, and how to migrate an existing Wasm light client contract.",source:"@site/docs/03-light-clients/04-wasm/05-governance.md",sourceDirName:"03-light-clients/04-wasm",slug:"/ibc/light-clients/wasm/governance",permalink:"/main/ibc/light-clients/wasm/governance",draft:!1,unlisted:!1,tags:[],version:"current",sidebarPosition:5,frontMatter:{title:"Governance",sidebar_label:"Governance",sidebar_position:5,slug:"/ibc/light-clients/wasm/governance"},sidebar:"defaultSidebar",previous:{title:"Messages",permalink:"/main/ibc/light-clients/wasm/messages"},next:{title:"Events",permalink:"/main/ibc/light-clients/wasm/events"}},r={},l=[{value:"Setting an authority",id:"setting-an-authority",level:2},{value:"Storing new Wasm light client byte code",id:"storing-new-wasm-light-client-byte-code",level:2},{value:"Migrating an existing Wasm light client contract",id:"migrating-an-existing-wasm-light-client-contract",level:2},{value:"Removing an existing checksum",id:"removing-an-existing-checksum",level:2}];function d(e){const t={a:"a",code:"code",h1:"h1",h2:"h2",p:"p",pre:"pre",...(0,o.a)(),...e.components};return(0,n.jsxs)(n.Fragment,{children:[(0,n.jsx)(t.h1,{id:"governance",children:"Governance"}),"\n",(0,n.jsx)(t.p,{children:"Learn how to upload Wasm light client byte code on a chain, and how to migrate an existing Wasm light client contract."}),"\n",(0,n.jsx)(t.h2,{id:"setting-an-authority",children:"Setting an authority"}),"\n",(0,n.jsxs)(t.p,{children:["Both the storage of Wasm light client byte code as well as the migration of an existing Wasm light client contract are permissioned (i.e. only allowed to an authority such as governance). The designated authority is specified when instantiating ",(0,n.jsx)(t.code,{children:"08-wasm"}),"'s keeper: both ",(0,n.jsx)(t.a,{href:"https://github.com/cosmos/ibc-go/blob/57fcdb9a9a9db9b206f7df2f955866dc4e10fef4/modules/light-clients/08-wasm/keeper/keeper.go#L39-L47",children:(0,n.jsx)(t.code,{children:"NewKeeperWithVM"})})," and ",(0,n.jsx)(t.a,{href:"https://github.com/cosmos/ibc-go/blob/57fcdb9a9a9db9b206f7df2f955866dc4e10fef4/modules/light-clients/08-wasm/keeper/keeper.go#L88-L96",children:(0,n.jsx)(t.code,{children:"NewKeeperWithConfig"})})," constructor functions accept an ",(0,n.jsx)(t.code,{children:"authority"})," argument that must be the address of the authorized actor. For example, in ",(0,n.jsx)(t.code,{children:"app.go"}),", when instantiating the keeper, you can pass the address of the governance module:"]}),"\n",(0,n.jsx)(t.pre,{children:(0,n.jsx)(t.code,{className:"language-go",children:'// app.go\nimport (\n  ...\n  "github.com/cosmos/cosmos-sdk/runtime"\n  authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"\n  govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"\n\n  ibcwasmkeeper "github.com/cosmos/ibc-go/modules/light-clients/08-wasm/keeper"\n  ibcwasmtypes "github.com/cosmos/ibc-go/modules/light-clients/08-wasm/types"\n  ...\n)\n\n// app.go\napp.WasmClientKeeper = ibcwasmkeeper.NewKeeperWithVM(\n  appCodec,\n  runtime.NewKVStoreService(keys[ibcwasmtypes.StoreKey]),\n  app.IBCKeeper.ClientKeeper,\n \tauthtypes.NewModuleAddress(govtypes.ModuleName).String(), // authority\n  wasmVM,\n  app.GRPCQueryRouter(),\n)\n'})}),"\n",(0,n.jsx)(t.h2,{id:"storing-new-wasm-light-client-byte-code",children:"Storing new Wasm light client byte code"}),"\n",(0,n.jsxs)(t.p,{children:["If governance is the allowed authority, the governance v1 proposal that needs to be submitted to upload a new light client contract should contain the message ",(0,n.jsx)(t.a,{href:"https://github.com/cosmos/ibc-go/blob/57fcdb9a9a9db9b206f7df2f955866dc4e10fef4/proto/ibc/lightclients/wasm/v1/tx.proto#L23-L30",children:(0,n.jsx)(t.code,{children:"MsgStoreCode"})})," with the base64-encoded byte code of the Wasm contract. Use the following CLI command and JSON as an example:"]}),"\n",(0,n.jsx)(t.pre,{children:(0,n.jsx)(t.code,{className:"language-shell",children:"simd tx gov submit-proposal <path/to/proposal.json> --from <key_or_address>\n"})}),"\n",(0,n.jsxs)(t.p,{children:["where ",(0,n.jsx)(t.code,{children:"proposal.json"})," contains:"]}),"\n",(0,n.jsx)(t.pre,{children:(0,n.jsx)(t.code,{className:"language-json",children:'{\n  "title": "Upload IBC Wasm light client",\n  "summary": "Upload wasm client",\n  "messages": [\n    {\n      "@type": "/ibc.lightclients.wasm.v1.MsgStoreCode",\n      "signer": "cosmos1...", // the authority address (e.g. the gov module account address)\n      "wasm_byte_code": "YWJ...PUB+" // standard base64 encoding of the Wasm contract byte code\n    }\n  ],\n  "metadata": "AQ==",\n  "deposit": "100stake"\n}\n'})}),"\n",(0,n.jsxs)(t.p,{children:["To learn more about the ",(0,n.jsx)(t.code,{children:"submit-proposal"})," CLI command, please check out ",(0,n.jsx)(t.a,{href:"https://docs.cosmos.network/main/modules/gov#submit-proposal",children:"the relevant section in Cosmos SDK documentation"}),"."]}),"\n",(0,n.jsxs)(t.p,{children:["Alternatively, the process of submitting the proposal may be simpler if you use the CLI command ",(0,n.jsx)(t.code,{children:"store-code"}),". This CLI command accepts as argument the file of the Wasm light client contract and takes care of constructing the proposal message with ",(0,n.jsx)(t.code,{children:"MsgStoreCode"})," and broadcasting it. See section ",(0,n.jsx)(t.a,{href:"/main/ibc/light-clients/wasm/client#store-code",children:(0,n.jsx)(t.code,{children:"store-code"})})," for more information."]}),"\n",(0,n.jsx)(t.h2,{id:"migrating-an-existing-wasm-light-client-contract",children:"Migrating an existing Wasm light client contract"}),"\n",(0,n.jsxs)(t.p,{children:["If governance is the allowed authority, the governance v1 proposal that needs to be submitted to migrate an existing new Wasm light client contract should contain the message ",(0,n.jsx)(t.a,{href:"https://github.com/cosmos/ibc-go/blob/57fcdb9a9a9db9b206f7df2f955866dc4e10fef4/proto/ibc/lightclients/wasm/v1/tx.proto#L52-L63",children:(0,n.jsx)(t.code,{children:"MsgMigrateContract"})})," with the checksum of the Wasm byte code to migrate to. Use the following CLI command and JSON as an example:"]}),"\n",(0,n.jsx)(t.pre,{children:(0,n.jsx)(t.code,{className:"language-shell",children:"simd tx gov submit-proposal <path/to/proposal.json> --from <key_or_address>\n"})}),"\n",(0,n.jsxs)(t.p,{children:["where ",(0,n.jsx)(t.code,{children:"proposal.json"})," contains:"]}),"\n",(0,n.jsx)(t.pre,{children:(0,n.jsx)(t.code,{className:"language-json",children:'{\n  "title": "Migrate IBC Wasm light client",\n  "summary": "Migrate wasm client",\n  "messages": [\n    {\n      "@type": "/ibc.lightclients.wasm.v1.MsgMigrateContract",\n      "signer": "cosmos1...", // the authority address (e.g. the gov module account address)\n      "client_id": "08-wasm-1", // client identifier of the Wasm light client contract that will be migrated\n      "checksum": "a8ad...4dc0", // SHA-256 hash of the Wasm byte code to migrate to, previously stored with MsgStoreCode\n      "msg": "{}" // JSON-encoded message to be passed to the contract on migration\n    }\n  ],\n  "metadata": "AQ==",\n  "deposit": "100stake"\n}\n'})}),"\n",(0,n.jsxs)(t.p,{children:["To learn more about the ",(0,n.jsx)(t.code,{children:"submit-proposal"})," CLI command, please check out ",(0,n.jsx)(t.a,{href:"https://docs.cosmos.network/main/modules/gov#submit-proposal",children:"the relevant section in Cosmos SDK documentation"}),"."]}),"\n",(0,n.jsx)(t.h2,{id:"removing-an-existing-checksum",children:"Removing an existing checksum"}),"\n",(0,n.jsxs)(t.p,{children:["If governance is the allowed authority, the governance v1 proposal that needs to be submitted to remove a specific checksum from the list of allowed checksums should contain the message ",(0,n.jsx)(t.a,{href:"https://github.com/cosmos/ibc-go/blob/57fcdb9a9a9db9b206f7df2f955866dc4e10fef4/proto/ibc/lightclients/wasm/v1/tx.proto#L39-L46",children:(0,n.jsx)(t.code,{children:"MsgRemoveChecksum"})})," with the checksum (of a corresponding Wasm byte code). Use the following CLI command and JSON as an example:"]}),"\n",(0,n.jsx)(t.pre,{children:(0,n.jsx)(t.code,{className:"language-shell",children:"simd tx gov submit-proposal <path/to/proposal.json> --from <key_or_address>\n"})}),"\n",(0,n.jsxs)(t.p,{children:["where ",(0,n.jsx)(t.code,{children:"proposal.json"})," contains:"]}),"\n",(0,n.jsx)(t.pre,{children:(0,n.jsx)(t.code,{className:"language-json",children:'{\n  "title": "Remove checksum of Wasm light client byte code",\n  "summary": "Remove checksum",\n  "messages": [\n    {\n      "@type": "/ibc.lightclients.wasm.v1.MsgRemoveChecksum",\n      "signer": "cosmos1...", // the authority address (e.g. the gov module account address)\n      "checksum": "a8ad...4dc0", // SHA-256 hash of the Wasm byte code that should be removed from the list of allowed checksums\n    }\n  ],\n  "metadata": "AQ==",\n  "deposit": "100stake"\n}\n'})}),"\n",(0,n.jsxs)(t.p,{children:["To learn more about the ",(0,n.jsx)(t.code,{children:"submit-proposal"})," CLI command, please check out ",(0,n.jsx)(t.a,{href:"https://docs.cosmos.network/main/modules/gov#submit-proposal",children:"the relevant section in Cosmos SDK documentation"}),"."]})]})}function h(e={}){const{wrapper:t}={...(0,o.a)(),...e.components};return t?(0,n.jsx)(t,{...e,children:(0,n.jsx)(d,{...e})}):d(e)}},11151:(e,t,s)=>{s.d(t,{Z:()=>c,a:()=>i});var n=s(67294);const o={},a=n.createContext(o);function i(e){const t=n.useContext(a);return n.useMemo((function(){return"function"==typeof e?e(t):{...t,...e}}),[t,e])}function c(e){let t;return t=e.disableParentContext?"function"==typeof e.components?e.components(o):e.components||o:i(e.components),n.createElement(a.Provider,{value:t},e.children)}}}]);