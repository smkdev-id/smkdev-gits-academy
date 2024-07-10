<script>
    export let y;

    export let tabs = [
        { name: "Learn", link: "#learn", dropdown: [
            { name: "Academy", link: "#academy" },
            { name: "Bootcamp", link: "#bootcamp" }
        ]},
        { name: "Community", link: "#community" },
        { name: "Blog", link: "https://www.youtube.com/watch?v=dQw4w9WgXcQ" },
        { name: "Dashboard", link: "#Login" },
    ];

    let showLearnDropdown = false;
    let showMobileMenu = false;
    let showMobileLearnDropdown = false;
    let hideDropdownTimeout;

    function showDropdown() {
        clearTimeout(hideDropdownTimeout);
        showLearnDropdown = true;
    }

    function hideDropdown() {
        hideDropdownTimeout = setTimeout(() => {
            showLearnDropdown = false;
        }, 200); // Adjust the delay as needed
    }

    function toggleMobileMenu() {
        showMobileMenu = !showMobileMenu;
    }

    function toggleMobileLearnDropdown() {
        showMobileLearnDropdown = !showMobileLearnDropdown;
    }
</script>

<style>
    @media (max-width: 640px) {
        .desktop-menu {
            display: none;
        }
        .mobile-menu {
            display: flex;
        }
        .mobile-nav {
            display: block;
        }
    }
    @media (min-width: 641px) {
        .desktop-menu {
            display: flex;
        }
        .mobile-menu {
            display: none;
        }
        .mobile-nav {
            display: none;
        }
    }
</style>

<header
    class={"sticky z-[10] top-0 duration-200 px-6 lg:px-[300px] flex items-center justify-between border-b border-solid " +
        (y > 0
            ? " py-2 bg-white border-violet-950"
            : " py-4 bg-transparent border-transparent")}
>
    <h1 class="font-medium">
        <a href="/"><img src="/images/logo-smkdev.png" width="110" height="23" alt="smkdev"></a>
    </h1>
    <div class="desktop-menu items-center gap-4">
        {#each tabs as tab, index}
        {#if tab.name == "Learn"}
            <button 
                class="relative"
                on:mouseenter={showDropdown}
                on:mouseleave={hideDropdown}
                aria-haspopup="true"
                aria-expanded={showLearnDropdown}
            >
                <a
                    href={tab.link}
                    class="duration-200 hover:text-violet-400 cursor-pointer"
                >
                    <p class="text-indigo-700">{tab.name}</p>
                </a>
                {#if showLearnDropdown}
                    <div class="absolute top-full mt-2 bg-white shadow-lg rounded-md">
                        {#each tab.dropdown as item}
                            <a href={item.link} class="block  px-4 py-2 text-indigo-700 hover:bg-violet-100">
                                {item.name}
                            </a>
                        {/each}
                    </div>
                {/if}
            </button>
        {/if}
        {#if tab.name != "Dashboard" && tab.name != "Learn"}
            <a
                href={tab.link}
                class="duration-200 hover:text-violet-400"
                target={index === 3 ? "_blank" : ""}
            >
                <p class="text-indigo-700">{tab.name}</p>
            </a>
        {/if}
        {/each}
        {#each tabs as tab}
        {#if tab.name == "Dashboard"}
        <a href={tab.link} class="blueShadow relative overflow-hidden px-5 py-2 group rounded-md bg-violet-900 text-slate-950">
            <div
                class="absolute top-0 right-full w-full h-full bg-violet-400 opacity-20 group-hover:translate-x-full z-0 duration-200"
            />
            <h4 class="relative z-9 text-white">{tab.name}</h4>
        </a>
        {/if}
        {/each}
    </div>
    <div class="mobile-menu">
        <button
            class="hamburger"
            on:click={toggleMobileMenu}
            aria-label="Toggle navigation"
        >
            {showMobileMenu ? '✖' : '☰'}
        </button>
    </div>
</header>

{#if showMobileMenu}
<nav class="mobile-nav bg-white shadow-lg rounded-md mt-2 px-6 py-4 z-[9]">
    {#each tabs as tab, index}
        <div class="relative mb-2">
            {#if tab.name == "Learn"}
                <a href="#/" class="block px-4 py-2 text-indigo-700 hover:bg-violet-100" on:click={toggleMobileLearnDropdown}>
                    {tab.name}
                </a>
                {#if showMobileLearnDropdown}
                    <div class="ml-4">
                        {#each tab.dropdown as item}
                            <a href={item.link} class="block px-4 py-2 text-indigo-700 hover:bg-violet-100">
                                {item.name}
                            </a>
                        {/each}
                    </div>
                {/if}
            {:else}
                <a
                    href={tab.link}
                    class="block px-4 py-2 text-indigo-700 hover:bg-violet-100"
                    target={index === 3 ? "_blank" : ""}
                >
                    {tab.name}
                </a>
            {/if}
        </div>
    {/each}
</nav>
{/if}
