import React, { useState, useEffect } from "react";
import CountdownTime from "../components/countdown/CountdownTime";
import Spacer from "../components/Spacer";

const LayoutCountdown = () => {
  const padding =
    "pt-[45px] pb-[75px] md:pt-[0px] md:pb-[135px] px-[30px] md:px-[135px]";
  const bgLight = "bg-gradient-to-r from-bgHeroLight-start via-bgHeroLight to-bgHeroLight-end";
  const bgDark = "dark:from-gray-900 dark:via-gray-800 dark:to-black";

  const [seconds, setSeconds] = useState(3);
  const [minutes, setMinutes] = useState(59);
  const [hours, setHours] = useState(23);
  const [days, setDays] = useState(15);

  useEffect(() => {
    const timer = setInterval(() => {
      setSeconds((prevSeconds) => {
        if (prevSeconds === 0) {
          setMinutes((prevMinutes) => {
            if (prevMinutes === 0) {
              setHours((prevHours) => {
                if (prevHours === 0) {
                  setDays((prevDays) => (prevDays === 0 ? 15 : prevDays - 1));
                  return 23;
                } else {
                  return prevHours - 1;
                }
              });
              return 59;
            } else {
              return prevMinutes - 1;
            }
          });
          return 59;
        } else {
          return prevSeconds - 1;
        }
      });
    }, 1000);

    return () => clearInterval(timer);
  }, []);

  return (
    <section className={`${bgLight} ${bgDark} ${padding}`}>
      <div className="w-full bg-orange-50 h-56 rounded-3xl overflow-hidden md:h-80">
        <img src="/src/assets/promo.jpeg" alt="manpromo" className="w-full h-full" />
      </div>

      <Spacer height={"h-5"} />

      <div className="flex justify-center">
        <CountdownTime
          days={days}
          hours={hours}
          minutes={minutes}
          seconds={seconds}
        />
      </div>
    </section>
  );
};

export default LayoutCountdown;
