import Link from "next/link";

export default function Home() {
  return (
    <div className="relative h-screen">
      {/* Video Background */}
      <video
        autoPlay
        loop
        muted
        playsInline
        className="absolute z-10 w-full h-full object-cover"
      >
        <source src="/video/toothVideo.mp4" type="video/mp4" />
      </video>

      {/* Overlay content */}
      <div className="relative z-20 flex justify-center items-center h-full">
        <Link href="/main">
          <div className="text-black text-lg font-bold bg-white w-[100px] h-[50px] text-center items-center flex justify-center content-center rounded-2xl">
            START
          </div>
        </Link>
      </div>
    </div>
  );
}
