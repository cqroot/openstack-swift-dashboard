import request from "./request";

export const getTargets = () => request.get("/target");
export const getDisks = (target) => request.get("/disk/" + target.ID);
