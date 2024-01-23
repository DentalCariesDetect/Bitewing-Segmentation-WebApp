import React from "react";
import { motion, useMotionValue, useSpring, useTransform } from "framer-motion";

// props type
type HoverCardProps = {
  title: string;
  message: string;
};

type TiltCardProps = {
  title: string;
  message: string;
};

const HoverCard: React.FC<HoverCardProps> = ({ title, message }) => {
  return (
    <div className="grid w-full place-content-center text-slate-900">
      <TiltCard title={title} message={message} />
    </div>
  );
};

const TiltCard: React.FC<TiltCardProps> = ({ title, message }) => {
  const x = useMotionValue(0);
  const y = useMotionValue(0);

  const mouseXSpring = useSpring(x);
  const mouseYSpring = useSpring(y);

  const rotateX = useTransform(
    mouseYSpring,
    [-0.5, 0.5],
    ["17.5deg", "-17.5deg"]
  );
  const rotateY = useTransform(
    mouseXSpring,
    [-0.5, 0.5],
    ["-17.5deg", "17.5deg"]
  );

  const handleMouseMove = (e: React.MouseEvent<HTMLDivElement, MouseEvent>) => {
    const rect = e.currentTarget.getBoundingClientRect();

    const width = rect.width;
    const height = rect.height;

    const mouseX = e.clientX - rect.left;
    const mouseY = e.clientY - rect.top;

    const xPct = mouseX / width - 0.5;
    const yPct = mouseY / height - 0.5;

    x.set(xPct);
    y.set(yPct);
  };

  const handleMouseLeave = () => {
    x.set(0);
    y.set(0);
  };

  return (
    <motion.div
      onMouseMove={handleMouseMove}
      onMouseLeave={handleMouseLeave}
      style={{
        rotateY,
        rotateX,
        transformStyle: "preserve-3d",
      }}
      className="relative h-96 w-[430px] rounded-xl bg-gradient-to-br from-indigo-300 to-violet-300"
    >
      <div
        style={{
          transform: "translateZ(75px)",
          transformStyle: "preserve-3d",
        }}
        className="absolute inset-4 grid place-content-center rounded-xl bg-white shadow-lg"
      >
        <p
          style={{
            transform: "translateZ(50px)",
          }}
          className="text-center text-2xl "
        >
          {title}
        </p>
        <div className="text-black text-xl font-light leading-normal text-shadow m-5">
          {message}
        </div>
      </div>
    </motion.div>
  );
};

export default HoverCard;
