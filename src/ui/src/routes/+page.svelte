<script lang="ts">
  import Navigator from "../components/common/navbar/Navigator.svelte";
  import { Button } from "$lib/components/ui/button";
  import { t } from "$lib/i18n";
  import { ROUTES } from "$lib/routes";
  import titlePageTop from "$lib/assets/title-page-top.png";
  import howToSetupPic from "$lib/assets/how-to-setup-pic.png";
  import {
    Gear,
    Target,
    Dashboard,
    Checkbox,
    PlusCircled,
    CodesandboxLogo,
    MagicWand,
  } from "svelte-radix";
  import * as Avatar from "$lib/components/ui/avatar";
  import * as Accordion from "$lib/components/ui/accordion";
  import Footer from "../components/common/footer/Footer.svelte";

  const navigationItems = [
    {
      key: "common.navigation.scroll-to-functions",
      action: () => {
        const featuresSection = document.getElementById("features");
        if (featuresSection) {
          featuresSection.scrollIntoView({ behavior: "smooth" });
        }
      },
    },
    {
      key: "common.navigation.scroll-to-step-by-step",
      action: () => {
        const stepsSection = document.getElementById("steps");
        if (stepsSection) {
          stepsSection.scrollIntoView({ behavior: "smooth" });
        }
      },
    },
    {
      key: "common.navigation.scroll-to-reviews",
      action: () => {
        const stepsSection = document.getElementById("reviews");
        if (stepsSection) {
          stepsSection.scrollIntoView({ behavior: "smooth" });
        }
      },
    },
    {
      key: "common.navigation.scroll-to-faq",
      action: () => {
        const faqSection = document.getElementById("faq");
        if (faqSection) {
          faqSection.scrollIntoView({ behavior: "smooth" });
        }
      },
    },
  ];

  const features = [
    {
      icon: Gear,
      title: "home.features.fast-setup",
      description: "home.features.fast-setup.description",
    },
    {
      icon: Target,
      title: "home.features.wide-target-audience",
      description: "home.features.wide-target-audience.description",
    },
    {
      icon: Dashboard,
      title: "home.features.overview-of-products",
      description: "home.features.overview-of-products.description",
    },
    {
      icon: Checkbox,
      title: "home.features.safe-payment",
      description: "home.features.safe-payment.description",
    },
  ];

  const steps = [
    {
      icon: PlusCircled,
      title: "home.steps.create-shop.title",
      description: "home.steps.create-shop.description",
    },
    {
      icon: CodesandboxLogo,
      title: "home.steps.add-products.title",
      description: "home.steps.add-products.description",
    },
    {
      icon: MagicWand,
      title: "home.steps.rent-products.title",
      description: "home.steps.rent-products.description",
    },
  ];

  const reviews = [
    {
      text: "home.reviews.review1.text",
      name: "home.reviews.review1.name",
      role: "home.reviews.review1.role",
      avatar: "home.reviews.review1.avatar",
    },
    {
      text: "home.reviews.review2.text",
      name: "home.reviews.review2.name",
      role: "home.reviews.review2.role",
      avatar: "home.reviews.review2.avatar",
    },
  ];

  const faqs = [
    {
      question: "home.faq.question1",
      answer: "home.faq.answer1",
    },
    {
      question: "home.faq.question2",
      answer: "home.faq.answer2",
    },
    {
      question: "home.faq.question3",
      answer: "home.faq.answer3",
    },
  ];
</script>

<Navigator {navigationItems} />

<main>
  <!-- Hero Section -->
  <div class="container mx-auto px-2">
    <div class="flex flex-col justify-between">
      <div class="grid grid-cols-1 md:grid-cols-2 gap-16 items-start pt-24">
        <!-- Left Column -->
        <div class="space-y-6">
          <h1 class="text-5xl font-bold tracking-tight">
            {$t("home.hero.title")}
          </h1>
        </div>

        <!-- Right Column -->
        <div class="space-y-8">
          <p class="text-lg text-muted-foreground">
            {$t("home.hero.description")}
          </p>
          <div class="flex flex-col sm:flex-row gap-4">
            <Button size="lg" href={ROUTES.SHOP} variant="default">
              {$t("home.hero.primary-cta")}
            </Button>
            <Button size="lg" href={ROUTES.SHOP} variant="link">
              {$t("home.hero.secondary-cta")}
            </Button>
          </div>
        </div>
      </div>

      <!-- Image Section -->
      <div class="w-full h-[50%] relative overflow-hidden rounded-lg pt-16">
        <img
          src={titlePageTop}
          alt="Title Page Top"
          class="w-full h-full object-cover"
        />
      </div>
    </div>
  </div>

  <!-- Features Section -->
  <section id="features" class="bg-muted/50 mt-16">
    <div class="container mx-auto px-12 py-12 pb-20">
      <div class="text-center mb-8">
        <p class="text-md text-muted-foreground mt-4 mb-6">
          {$t("home.features.subtitle")}
        </p>
        <h2 class="text-4xl font-semibold">
          {$t("home.features.title")}
        </h2>
        <p class="text-md text-muted-foreground mt-6 mb-12">
          {$t("home.features.description")}
        </p>
      </div>

      <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-8">
        {#each features as feature}
          <div class="flex flex-col items-center text-center space-y-4">
            <div class="p-3 bg-background rounded-xl border border-border">
              <svelte:component this={feature.icon} class="h-6 w-6" />
            </div>
            <h3 class="font-medium">{$t(feature.title)}</h3>
            <p class="text-sm text-muted-foreground">
              {$t(feature.description)}
            </p>
          </div>
        {/each}
      </div>
    </div>
  </section>

  <!-- Steps Section -->
  <section id="steps" class="container mx-auto px-12 py-20">
    <div class="grid grid-cols-1 lg:grid-cols-2 gap-16 items-center">
      <!-- Left Column -->
      <div class="space-y-12">
        <div class="space-y-6">
          <p class="text-md text-muted-foreground">
            {$t("home.steps.subtitle")}
          </p>
          <h2 class="text-4xl font-semibold">
            {$t("home.steps.title")}
          </h2>
          <p class="text-lg text-muted-foreground">
            {$t("home.steps.description")}
          </p>
        </div>

        <div class="space-y-8">
          {#each steps as step}
            <div class="flex gap-10">
              <svelte:component this={step.icon} class="h-8 w-8" />
              <div class="space-y-2">
                <h3 class="font-medium">{$t(step.title)}</h3>
                <p class="text-sm text-muted-foreground">
                  {$t(step.description)}
                </p>
              </div>
            </div>
          {/each}
        </div>
      </div>

      <!-- Right Column - Image -->
      <div class="relative rounded-lg overflow-hidden">
        <img
          src={howToSetupPic}
          alt="Step by Step Guide"
          class="w-full h-full object-cover"
        />
      </div>
    </div>
  </section>

  <!-- Reviews Section -->
  <section id="reviews" class="bg-muted/50 mt-16">
    <div class="container mx-auto px-12 py-20">
      <div class="text-left mb-16">
        <p class="text-md text-muted-foreground mb-4">
          {$t("home.reviews.subtitle")}
        </p>
        <h2 class="text-4xl font-semibold">
          {$t("home.reviews.title")}
        </h2>
      </div>

      <div class="grid grid-cols-1 md:grid-cols-2 gap-24">
        {#each reviews as review}
          <div class="rounded-xl space-y-8">
            <p class="text-lg">
              {$t(review.text)}
            </p>
            <div class="flex items-center gap-4">
              <Avatar.Root>
                <Avatar.Fallback>
                  {$t(review.avatar)}
                </Avatar.Fallback>
              </Avatar.Root>
              <div>
                <p class="font-medium">{$t(review.name)}</p>
                <p class="text-sm text-muted-foreground">{$t(review.role)}</p>
              </div>
            </div>
          </div>
        {/each}
      </div>
    </div>
  </section>

  <!-- FAQ Section -->
  <section id="faq" class="container mx-auto px-12 py-20">
    <div class="text-center mb-16">
      <p class="text-md text-muted-foreground mb-4">
        {$t("home.faq.subtitle")}
      </p>
      <h2 class="text-4xl font-semibold mb-6 whitespace-pre-wrap">
        {$t("home.faq.title")}
      </h2>
      <p class="text-lg text-muted-foreground max-w-2xl mx-auto">
        {$t("home.faq.description")}
      </p>
    </div>

    <div class="max-w-3xl mx-auto mb-16">
      <Accordion.Root>
        {#each faqs as faq}
          <Accordion.Item value={faq.question}>
            <Accordion.Trigger>
              {$t(faq.question)}
            </Accordion.Trigger>
            <Accordion.Content>
              {$t(faq.answer)}
            </Accordion.Content>
          </Accordion.Item>
        {/each}
      </Accordion.Root>
    </div>

    <!-- Contact Card -->
    <div class="max-w-2xl mx-auto bg-muted/50 rounded-xl p-8 text-center">
      <h3 class="text-2xl font-semibold mb-4">
        {$t("home.faq.contact.title")}
      </h3>
      <p class="text-muted-foreground mb-6">
        {$t("home.faq.contact.description")}
      </p>
      <Button href="mailto:support@example.com" variant="default">
        {$t("home.faq.contact.button")}
      </Button>
    </div>
  </section>
</main>

<Footer footerItems={navigationItems} />