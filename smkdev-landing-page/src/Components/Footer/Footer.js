import React from 'react';
import Logo from '../../Assets/Logo.png';
import { BsTwitter } from 'react-icons/bs';
import { SiLinkedin } from 'react-icons/si';
import { BsYoutube } from 'react-icons/bs';
import { FaFacebookF } from 'react-icons/fa';

const Footer = () => {
  return (
    <div className='footer-wrapper'>
      <div className='footer-content'>
        <div className='footer-logo-container'>
          <img src={Logo} alt='' width={200} />
          <p className='fs-5 mt-1 fw-bold'>
            Creating High-Caliber Digital Talent
          </p>
        </div>
        <div className='footer-description'>
          <p className='fs-6'>
            Summarecon Bandung, Jl. Magna Barat Blok MD No.02, Rancabolang, Kec.
            Gedebage, Kota Bandung, Jawa Barat 40295
          </p>
        </div>
        <div className='footer-icons-group'>
          <p className='fw-bold fs-5'>Follow Us</p>
          <div className='footer-icons d-flex gap-3'>
            <BsTwitter className='text-primary fs-3' />
            <SiLinkedin className='text-primary fs-3' />
            <BsYoutube className='text-primary fs-3' />
            <FaFacebookF className='text-primary fs-3' />
          </div>
        </div>
      </div>
      <div className='footer-copyright justify-content-center d-flex'>
        <p>
          &copy;2024 SMKDEV – PT Eureka Merdeka Indonesia. All Rights Reserved.
        </p>
      </div>
    </div>
  );
};

export default Footer;
