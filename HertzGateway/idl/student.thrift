namespace go gateway

struct College {
    1: required string name(go.tag = 'json:"name"'),
    2: string address(go.tag = 'json:"address"'),
}

struct Student {
    1: required i32 id(go.tag = 'json:"id"'),
    2: required string name(go.tag = 'json:"name"'),
    3: required College college(go.tag = 'json:"college"'),
    4: optional list<string> email(go.tag = 'json:"email"'),
}

struct BizRequest {
    1: optional Student student(api.body = 'student'),
    2: optional i32 itemId(api.body = 'itemId')
    3: optional string method(api.path = 'method')
    4: optional string service(api.path = 'service')
}

struct RspItem {
    1: optional i64 item_id
    2: optional string text
}

struct BizResponse {
    1: optional bool success(api.body='success'),
    2: optional string message(api.body='message'),
    3: optional Student student(api.body = 'student')
}

service BizService {
    BizResponse Register(1: BizRequest req)(api.post = '/RPC/:service/:method', api.param = 'true', api.serializer = 'json')
    BizResponse Query(1: BizRequest req)(api.get = '/RPC/:service/:method', api.param = 'true')
}
