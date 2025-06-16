import {
  type RouteConfig,
  route,
  index,
  prefix,
} from "@react-router/dev/routes";

export default [
  index("routes/home.tsx"),
  route("auth/callback", "routes/auth.callback.tsx"),
  route("logout", "routes/logout.tsx"),
  ...prefix("ratings", [
    index("routes/ratings/home.tsx"),
    route(":festivalName", "routes/ratings/festival.tsx"),
  ]),
  ...prefix("festivals", [route(":name", "routes/festivals/festival.tsx")]),
] satisfies RouteConfig;
