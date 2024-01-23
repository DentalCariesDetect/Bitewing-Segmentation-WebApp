import { Inter } from "next/font/google";

const inter = Inter({ subsets: ["latin"] });

export default function Home() {
  return (
    // Full screen background
    <div className="bg-gradient-background h-screen">
      <h1>Tutorial 3</h1>
    </div>
  );
}
