import { useEffect, useRef } from "react";
import { gsap } from "gsap";
import { ScrollTrigger } from "gsap/ScrollTrigger";

const TalentaDigital = () => {
  const imageRef = useRef(null);

  useEffect(() => {
    gsap.registerPlugin(ScrollTrigger);

    const imageElement = imageRef.current;

    gsap.set(imageElement, { opacity: 0 }); 

    ScrollTrigger.create({
      trigger: imageElement,
      start: "top 65%",
      end: "top 30%", 
      onUpdate: self => {
        if (self.direction === 1) {
          // Scroll down
          gsap.to(imageElement, {
            opacity: 1,
            duration: 4,
            ease: "power2.out",
          });
        }
      },
    });

    
    return () => {
      ScrollTrigger.getAll().forEach(trigger => trigger.kill());
    };
  }, []);

  return (
    <section className="talenta-digital-container relative h-screen px-24 overflow-hidden flex flex-col items-center justify-center bg-cover bg-center bg-talent-digital">
      <div className="relative z-10 flex flex-col items-center justify-center">
        <h1 className="text-3xl font-medium mb-4 text-center text-custom-blue">
          Talenta Digital SMKDEV
        </h1>
        <img
          ref={imageRef}
          src="./public/Professional-5-Steps-SMKDEV-Build-Digital-Talent-2.png"
          alt="Professional Steps"
          width={1000}
          className="max-w-full mx-auto"
        />
      </div>
    </section>
  );
};

export default TalentaDigital;
