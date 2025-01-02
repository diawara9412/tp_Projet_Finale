import type { AppType } from "next/app";
import "./globals.css";
import { trpc } from "@/trpc/trpc";

const MyApp: AppType = ({ Component, pageProps }) => {
  return <Component {...pageProps} />;
};
export default trpc.withTRPC(MyApp);
