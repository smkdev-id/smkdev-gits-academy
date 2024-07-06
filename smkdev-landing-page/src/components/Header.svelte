<script>
    export let y;

    export let tabs = [
        { name: "Learn ", link: "#learn" },
        { name: "Community", link: "#community" },
        { name: "Blog", link: "https://www.youtube.com/watch?v=dQw4w9WgXcQ" },
        { name: "Dashboard", link: "#Login" },
        // {name: '', link: ''},
    ];
</script>

<header
    class={"sticky z-[10] top-0 duration-200 px-6 flex items-center justify-between border-b border-solid " +
        (y > 0
            ? " py-4 bg-white border-violet-950"
            : " py-6 bg-transparent border-transparent")}
>
    <h1 class="font-medium">
        <a href="/"><img src="/images/logo-smkdev.png" width="110" height="23" alt=""></a>
    </h1>
    <div class="sm:flex items-center gap-4 hidden">
        {#each tabs as tab, index}
        {#if tab.name != "Dashboard"}
            <a
                href={tab.link}
                class="duration-200 hover:text-violet-400"
                target={index === 3 ? "_blank" : ""}
            >
                <p class="text-indigo-700">{tab.name}</p>
            </a>
        {/if}
        {/each}
        <!-- <button
            class="blueShadow relative overflow-hidden px-5 py-2 group rounded-full bg-violet-600 text-white"
        >
            <div
                class="absolute top-0 right-full w-full h-full bg-violet-400 opacity-20 group-hover:translate-x-full z-0 duration-200"
            />
            <h4 class="relative z-9">Login</h4>
        </button> -->
        {#each tabs as tab}
        {#if tab.name == "Dashboard"}
        <a href={tab.link} class="blueShadow relative overflow-hidden px-5 py-2 group rounded-full bg-indigo-700 text-slate-950">
            <div
                class="absolute top-0 right-full w-full h-full bg-violet-400 opacity-20 group-hover:translate-x-full z-0 duration-200"
            />
            <h4 class="relative z-9 text-white">{tab.name}</h4>
        </a>
        {/if}
        {/each}
    </div>
</header>