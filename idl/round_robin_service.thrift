include "base.thrift"

namespace go github.DexMus.round_robin_service

struct RoundQueryGameInfoReq {
    1: required string game;
    2: required string gamerID;
    3: required i64 points;
}

struct RoundQueryGameInfoResp {
    1: required string game;
    2: required string gamerID;
    3: required i64 points;
}

struct QueryGameInfoReq {
    1: required string game;
    2: required string gamerID;
    3: required i64 points;
}

struct QueryGameInfoResp {
    1: required string game;
    2: required string gamerID;
    3: required i64 points;
}

service RoundRobinService {
    RoundQueryGameInfoResp RoundQueryGameInfo(1: RoundQueryGameInfoReq req) (api.post="/api/game_center/v1/meta/round_query/")
    QueryGameInfoResp QueryGameInfo(1: QueryGameInfoReq req) (api.post="/api/game_center/v1/meta/query/")
}
