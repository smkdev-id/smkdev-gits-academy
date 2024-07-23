import { TypeAnimation } from 'react-type-animation';

export const TextAnimation = () => {
  return (
    <TypeAnimation
      sequence={[
        'Global',
        2000,
        'Masa Depan Indonesia',
        2000,
      ]}
      wrapper="span"
      speed={2}
      repeat={Infinity}
    />
  );
};