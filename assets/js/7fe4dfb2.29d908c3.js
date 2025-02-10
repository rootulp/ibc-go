"use strict";(self.webpackChunkdocs=self.webpackChunkdocs||[]).push([[96959],{96367:(e,s,i)=>{i.r(s),i.d(s,{assets:()=>a,contentTitle:()=>r,default:()=>h,frontMatter:()=>t,metadata:()=>c,toc:()=>l});var n=i(85893),o=i(11151);const t={title:"IBC-Go v7.2 to v7.3",sidebar_label:"IBC-Go v7.2 to v7.3",sidebar_position:10,slug:"/migrations/v7_2-to-v7_3"},r="Migrating from v7.2 to v7.3",c={id:"migrations/v7_2-to-v7_3",title:"IBC-Go v7.2 to v7.3",description:"This guide provides instructions for migrating to version v7.3.0 of ibc-go.",source:"@site/versioned_docs/version-v9.0.x/05-migrations/10-v7_2-to-v7_3.md",sourceDirName:"05-migrations",slug:"/migrations/v7_2-to-v7_3",permalink:"/v9/migrations/v7_2-to-v7_3",draft:!1,unlisted:!1,tags:[],version:"v9.0.x",sidebarPosition:10,frontMatter:{title:"IBC-Go v7.2 to v7.3",sidebar_label:"IBC-Go v7.2 to v7.3",sidebar_position:10,slug:"/migrations/v7_2-to-v7_3"},sidebar:"defaultSidebar",previous:{title:"IBC-Go v7 to v7.1",permalink:"/v9/migrations/v7-to-v7_1"},next:{title:"IBC-Go v7 to v8",permalink:"/v9/migrations/v7-to-v8"}},a={},l=[{value:"Chains",id:"chains",level:2},{value:"IBC Apps",id:"ibc-apps",level:2},{value:"Relayers",id:"relayers",level:2},{value:"IBC Light Clients",id:"ibc-light-clients",level:2},{value:"06-solomachine",id:"06-solomachine",level:3}];function d(e){const s={a:"a",code:"code",h1:"h1",h2:"h2",h3:"h3",li:"li",p:"p",strong:"strong",ul:"ul",...(0,o.a)(),...e.components};return(0,n.jsxs)(n.Fragment,{children:[(0,n.jsx)(s.h1,{id:"migrating-from-v72-to-v73",children:"Migrating from v7.2 to v7.3"}),"\n",(0,n.jsxs)(s.p,{children:["This guide provides instructions for migrating to version ",(0,n.jsx)(s.code,{children:"v7.3.0"})," of ibc-go."]}),"\n",(0,n.jsx)(s.p,{children:"There are four sections based on the four potential user groups of this document:"}),"\n",(0,n.jsxs)(s.ul,{children:["\n",(0,n.jsxs)(s.li,{children:[(0,n.jsx)(s.a,{href:"#migrating-from-v72-to-v73",children:"Migrating from v7.2 to v7.3"}),"\n",(0,n.jsxs)(s.ul,{children:["\n",(0,n.jsx)(s.li,{children:(0,n.jsx)(s.a,{href:"#chains",children:"Chains"})}),"\n",(0,n.jsx)(s.li,{children:(0,n.jsx)(s.a,{href:"#ibc-apps",children:"IBC Apps"})}),"\n",(0,n.jsx)(s.li,{children:(0,n.jsx)(s.a,{href:"#relayers",children:"Relayers"})}),"\n",(0,n.jsx)(s.li,{children:(0,n.jsx)(s.a,{href:"#ibc-light-clients",children:"IBC Light Clients"})}),"\n"]}),"\n"]}),"\n"]}),"\n",(0,n.jsxs)(s.p,{children:[(0,n.jsx)(s.strong,{children:"Note:"})," ibc-go supports golang semantic versioning and therefore all imports must be updated on major version releases."]}),"\n",(0,n.jsx)(s.h2,{id:"chains",children:"Chains"}),"\n",(0,n.jsxs)(s.ul,{children:["\n",(0,n.jsx)(s.li,{children:"No relevant changes were made in this release."}),"\n"]}),"\n",(0,n.jsx)(s.h2,{id:"ibc-apps",children:"IBC Apps"}),"\n",(0,n.jsxs)(s.p,{children:["A set of interfaces have been added that IBC applications may optionally implement. Developers interested in integrating their applications with the ",(0,n.jsx)(s.a,{href:"/v9/middleware/callbacks/overview",children:"callbacks middleware"})," should implement these interfaces so that the callbacks middleware can retrieve the desired callback addresses on the source and destination chains and execute actions on packet lifecycle events. The interfaces are ",(0,n.jsx)(s.a,{href:"https://github.com/cosmos/ibc-go/blob/v7.3.0-rc1/modules/core/05-port/types/module.go#L142-L147",children:(0,n.jsx)(s.code,{children:"PacketDataUnmarshaler"})}),", ",(0,n.jsx)(s.a,{href:"https://github.com/cosmos/ibc-go/blob/v7.3.0-rc1/modules/core/exported/packet.go#L43-L52",children:(0,n.jsx)(s.code,{children:"PacketDataProvider"})})," and ",(0,n.jsx)(s.a,{href:"https://github.com/cosmos/ibc-go/blob/v7.3.0-rc1/modules/core/exported/packet.go#L36-L41",children:(0,n.jsx)(s.code,{children:"PacketData"})}),"."]}),"\n",(0,n.jsxs)(s.p,{children:["Sample implementations are available for reference. For ",(0,n.jsx)(s.code,{children:"transfer"}),":"]}),"\n",(0,n.jsxs)(s.ul,{children:["\n",(0,n.jsxs)(s.li,{children:[(0,n.jsx)(s.a,{href:"https://github.com/cosmos/ibc-go/blob/v7.3.0-rc1/modules/apps/transfer/ibc_module.go#L303-L313",children:(0,n.jsx)(s.code,{children:"PacketDataUnmarshaler"})}),","]}),"\n",(0,n.jsx)(s.li,{children:(0,n.jsx)(s.a,{href:"https://github.com/cosmos/ibc-go/blob/v7.3.0-rc1/modules/apps/transfer/types/packet.go#L85-L105",children:(0,n.jsx)(s.code,{children:"PacketDataProvider"})})}),"\n",(0,n.jsxs)(s.li,{children:["and ",(0,n.jsx)(s.a,{href:"https://github.com/cosmos/ibc-go/blob/v7.3.0-rc1/modules/apps/transfer/types/packet.go#L74-L83",children:(0,n.jsx)(s.code,{children:"PacketData"})}),"."]}),"\n"]}),"\n",(0,n.jsxs)(s.p,{children:["For ",(0,n.jsx)(s.code,{children:"27-interchain-accounts"}),":"]}),"\n",(0,n.jsxs)(s.ul,{children:["\n",(0,n.jsxs)(s.li,{children:[(0,n.jsx)(s.a,{href:"https://github.com/cosmos/ibc-go/blob/v7.3.0-rc1/modules/apps/27-interchain-accounts/controller/ibc_middleware.go#L258-L268",children:(0,n.jsx)(s.code,{children:"PacketDataUnmarshaler"})}),","]}),"\n",(0,n.jsx)(s.li,{children:(0,n.jsx)(s.a,{href:"https://github.com/cosmos/ibc-go/blob/v7.3.0-rc1/modules/apps/27-interchain-accounts/types/packet.go#L94-L114",children:(0,n.jsx)(s.code,{children:"PacketDataProvider"})})}),"\n",(0,n.jsxs)(s.li,{children:["and ",(0,n.jsx)(s.a,{href:"https://github.com/cosmos/ibc-go/blob/v7.3.0-rc1/modules/apps/27-interchain-accounts/types/packet.go#L78-L92",children:(0,n.jsx)(s.code,{children:"PacketData"})}),"."]}),"\n"]}),"\n",(0,n.jsx)(s.h2,{id:"relayers",children:"Relayers"}),"\n",(0,n.jsxs)(s.ul,{children:["\n",(0,n.jsx)(s.li,{children:"No relevant changes were made in this release."}),"\n"]}),"\n",(0,n.jsx)(s.h2,{id:"ibc-light-clients",children:"IBC Light Clients"}),"\n",(0,n.jsx)(s.h3,{id:"06-solomachine",children:"06-solomachine"}),"\n",(0,n.jsxs)(s.p,{children:["Solo machines are now expected to sign data on a path that 1) does not include a connection prefix (e.g ",(0,n.jsx)(s.code,{children:"ibc"}),") and 2) does not escape any characters. See PR ",(0,n.jsx)(s.a,{href:"https://github.com/cosmos/ibc-go/pull/4429",children:"#4429"})," for more details. We recommend ",(0,n.jsx)(s.strong,{children:"NOT"})," using the solo machine light client of versions lower than v7.3.0."]})]})}function h(e={}){const{wrapper:s}={...(0,o.a)(),...e.components};return s?(0,n.jsx)(s,{...e,children:(0,n.jsx)(d,{...e})}):d(e)}},11151:(e,s,i)=>{i.d(s,{Z:()=>c,a:()=>r});var n=i(67294);const o={},t=n.createContext(o);function r(e){const s=n.useContext(t);return n.useMemo((function(){return"function"==typeof e?e(s):{...s,...e}}),[s,e])}function c(e){let s;return s=e.disableParentContext?"function"==typeof e.components?e.components(o):e.components||o:r(e.components),n.createElement(t.Provider,{value:s},e.children)}}}]);