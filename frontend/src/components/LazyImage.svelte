<script lang="ts">
  export let src: string;
  export let alt = "";
  export let className = "";

  let loaded = false;
  let isVisible = false;

  function lazyload(node: HTMLElement) {
    let observer: IntersectionObserver;

    const onIntersect = (entries: IntersectionObserverEntry[]) => {
      const { isIntersecting } = entries[0];
      if (isIntersecting) {
        isVisible = true;
        observer.disconnect();
      }
    };

    observer = new IntersectionObserver(onIntersect, {
      rootMargin: "100px",
      threshold: 0.01,
    });

    observer.observe(node);

    return {
      destroy() {
        if (observer) {
          observer.disconnect();
        }
      },
    };
  }

  // Function to reload the image
  export function reloadImage() {
    loaded = false;

    const newSrc = `${src}?${Date.now()}`;

    const img = new Image();
    img.onload = () => {
      loaded = true;
    };
    img.src = newSrc;
    src = newSrc;
    isVisible = true; // Force visibility to trigger the image load
  }

  $: if (isVisible) {
    const img = new Image();
    img.onload = () => {
      loaded = true;
    };
    img.src = src;
  }
</script>

<div use:lazyload class="relative {className}">
  {#if !loaded && isVisible}
    <div class="absolute inset-0 flex items-center justify-center bg-gray-100">
      <div
        class="w-5 h-5 border-2 border-gray-300 border-t-blue-500 rounded-full animate-spin"
      ></div>
    </div>
  {/if}

  <img
    src={isVisible ? src : ""}
    {alt}
    class=" object-cover transition-opacity duration-300 {loaded
      ? 'opacity-100'
      : 'opacity-0'}"
  />
</div>
