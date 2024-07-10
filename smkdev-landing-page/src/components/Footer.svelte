<script>
    let footer = [
        {
            imgIcon : "/images/logo-smkdev.png",
            qoutes : "Creating High-Caliber Digital Talent",
            address : "Summarecon Bandung, Jl. Magna Barat Blok MD No.02, Rancabolang, Kec. Gedebage, Kota Bandung, Jawa Barat 40295",
            links : [
                {
                    name : "twitter",
                    link : "#",
                    icon : "fa-brands fa-square-x-twitter"
                },
                {
                    name : "Instagram",
                    link : "#",
                    icon : "fa-brands fa-instagram"
                },
                {
                    name : "facebook",
                    link : "#",
                    icon : "fa-brands fa-square-facebook"
                }

            ]
        }
    ]
</script>
<section class="relative pt-24 overflow-hidden px-6 lg:px-[300px]">
    <div class="relative z-10 container px-4 mx-auto">
      <div class="pb-10 border-b border-gray-900">
        <div class="flex flex-wrap -m-8">
          <div class="w-full sm:w-1/2 lg:w-full pb-6 px-6">
            {#each footer as item}
                 <!-- content here -->
                 <div class="lg:max-w-sm">
                   <a class="mb-4 inline-block" href="/">
                     <img src="{item.imgIcon}" alt="smkdev" width="110" height="23">
                   </a>
                   <p class="mb-4 text-md leading-relaxed md:max-w-xs">{item.qoutes}</p>
                   <p class="mb-5 font-medium leading-relaxed md:max-w-xs">{item.address}</p>
                   <div class="flex flex-wrap ">
                   {#each item.links as data}
                      <!-- content here -->
                        <div class="w-auto p-1.5">
                          <a href="{data.link}">
                            <div class="flex items-center justify-center w-9 h-9 rounded-full bg-indigo-100">
                              <i class="{data.icon}"></i>
                            </div>
                          </a>
                        </div>
                        {/each}
                    </div>
                 </div>
            {/each}
          </div>
      </div>
      <p class="mt-10 text-md text-center">Copyright © 2022 SMKDEV – PT Eureka Merdeka Indonesia. All Rights Reserved.</p>
    </div>
  </section>