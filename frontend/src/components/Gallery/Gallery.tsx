import { Swiper, SwiperSlide, useSwiper } from 'swiper/react';
import styles from './Gallery.module.scss';
import cn from 'classnames';
import { Fancybox } from '@fancyapps/ui';
import '@fancyapps/ui/dist/fancybox/fancybox.css';
import { Navigation } from 'swiper/modules';

import 'swiper/css';

interface IGalleryProps {
  images: string[];
}

Fancybox.bind('[data-fancybox="gallery"]');

const SwiperButtonNext = () => {
  const swiper = useSwiper();
  return (
    <button className={cn(styles.navButton)} onClick={() => swiper.slideNext()}>
      →
    </button>
  );
};

const SwiperButtonPrev = () => {
  const swiper = useSwiper();
  return (
    <button className={cn(styles.navButton, styles.reverse)} onClick={() => swiper.slidePrev()}>
      ←
    </button>
  );
};

export const Gallery = ({ images }: IGalleryProps) => {
  return (
    <div className={styles.sliderWrapper}>
      <Swiper
        modules={[Navigation]}
        slidesPerView={1}
        navigation
        className={cn('mySwiper', styles.centered)}
      >
        {images.map((item, index) => {
          return (
            <SwiperSlide key={index}>
              <div className={styles.itemWrapper}>
                <a href={item} data-fancybox="gallery">
                  <img src={item} />
                </a>
              </div>
            </SwiperSlide>
          );
        })}
        <SwiperButtonPrev />
        <SwiperButtonNext />
      </Swiper>
    </div>
  );
};
