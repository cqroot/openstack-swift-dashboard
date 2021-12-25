const drawerRouter = [
  {
    path: "/dashboard",
    name: "Dashboard",
    component: () => import("@/views/Dashboard.vue"),
    meta: { icon: "mdi-view-dashboard" },
  },
  {
    path: "/disks",
    name: "Disks",
    component: () => import("@/views/Disks.vue"),
    meta: { icon: "mdi-harddisk" },
  },
];

export default drawerRouter;
