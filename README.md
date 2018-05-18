## JSON to shell variables converter

`js2sh` is a simple tool to convert _JSON_ data to shell variables.

### Usage

Piping content to `js2sh`

```
$ echo '{"foo": [1, 2, 3, 8.923878], "bar": {"baz": {"a": "b"}}}' | js2sh
FOO_0="1"
FOO_1="2"
FOO_2="3"
FOO_3="8.923878"
BAR_BAZ_A="b"
```

Passing a _JSON_ file
```
$ js2sh /path/to/foo.json
[...]
RESOURCES_VPC_LOAD_REQUEST_OPERATION="DescribeVpcs"
RESOURCES_VPC_LOAD_REQUEST_PARAMS_0_SOURCE="identifier"
RESOURCES_VPC_LOAD_REQUEST_PARAMS_0_NAME="Id"
RESOURCES_VPC_LOAD_REQUEST_PARAMS_0_TARGET="VpcIds[0]"
RESOURCES_VPC_LOAD_PATH="Vpcs[0]"
RESOURCES_VPC_ACTIONS_CREATENETWORKACL_REQUEST_OPERATION="CreateNetworkAcl"
RESOURCES_VPC_ACTIONS_CREATENETWORKACL_REQUEST_PARAMS_0_NAME="Id"
RESOURCES_VPC_ACTIONS_CREATENETWORKACL_REQUEST_PARAMS_0_TARGET="VpcId"
RESOURCES_VPC_ACTIONS_CREATENETWORKACL_REQUEST_PARAMS_0_SOURCE="identifier"
RESOURCES_VPC_ACTIONS_CREATENETWORKACL_RESOURCE_TYPE="NetworkAcl"
RESOURCES_VPC_ACTIONS_CREATENETWORKACL_RESOURCE_IDENTIFIERS_0_TARGET="Id"
RESOURCES_VPC_ACTIONS_CREATENETWORKACL_RESOURCE_IDENTIFIERS_0_SOURCE="response"
RESOURCES_VPC_ACTIONS_CREATENETWORKACL_RESOURCE_IDENTIFIERS_0_PATH="NetworkAcl.NetworkAclId"
[…]
```

Do not upper-case variables

```
$ js2sh -n example.json
glossary_title="example glossary"
glossary_GlossDiv_title="S"
glossary_GlossDiv_GlossList_GlossEntry_Abbrev="ISO 8879:1986"
glossary_GlossDiv_GlossList_GlossEntry_GlossDef_GlossSeeAlso_0="GML"
glossary_GlossDiv_GlossList_GlossEntry_GlossDef_GlossSeeAlso_1="XML"
glossary_GlossDiv_GlossList_GlossEntry_GlossDef_para="A meta-markup language, used to create markup languages such as DocBook."
glossary_GlossDiv_GlossList_GlossEntry_GlossSee="markup"
glossary_GlossDiv_GlossList_GlossEntry_ID="SGML"
glossary_GlossDiv_GlossList_GlossEntry_SortAs="SGML"
glossary_GlossDiv_GlossList_GlossEntry_GlossTerm="Standard Generalized Markup Language"
glossary_GlossDiv_GlossList_GlossEntry_Acronym="SGML"
```

Filter matching values
```
$ js2sh -f TITLE example.json
GLOSSARY_TITLE="example glossary"
GLOSSARY_GLOSSDIV_TITLE="S"
```

Modify default separator

```
$ js2sh -s @@ example.json
GLOSSARY@@TITLE="example glossary"
GLOSSARY@@GLOSSDIV@@GLOSSLIST@@GLOSSENTRY@@GLOSSTERM="Standard Generalized Markup Language"
GLOSSARY@@GLOSSDIV@@GLOSSLIST@@GLOSSENTRY@@ACRONYM="SGML"
GLOSSARY@@GLOSSDIV@@GLOSSLIST@@GLOSSENTRY@@ABBREV="ISO 8879:1986"
GLOSSARY@@GLOSSDIV@@GLOSSLIST@@GLOSSENTRY@@GLOSSDEF@@PARA="A meta-markup language, used to create markup languages such as DocBook."
```
