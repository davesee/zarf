<script lang="ts">
	import AccordionGroup from '../../../lib/components/accordion-group.svelte';

	import Icon from '$lib/components/icon.svelte';
	import PackageDetails from '$lib/components/package-details-card.svelte';
	import PackageComponent from '$lib/components/package-component-accordion.svelte';
	import { pkgStore } from '$lib/store';
	import { Button, Typography } from '@ui';
</script>

<svelte:head>
	<title>Configure</title>
</svelte:head>
<section class="page-header">
	<Typography variant="h4">Configure Package Deployment</Typography>
</section>

<section class="page-section">
	<Typography variant="h5">
		<Icon variant="package" />
		Package Details
	</Typography>
	<PackageDetails pkg={$pkgStore.zarfPackage} />
</section>

<section class="page-section">
	<Typography variant="h5">
		<Icon variant="component" />
		Package Components
		<Typography variant="caption" element="p">
			<span aria-hidden="true">
				<Icon variant="component" className="invisible" />
			</span>
			The following components will be deployed into the cluster. Optional components that are not selected
			will not be deployed.
		</Typography>
	</Typography>

	<AccordionGroup>
		{#each $pkgStore.zarfPackage.components as component, idx}
			<PackageComponent {idx} {component} readOnly={false} />
		{/each}
	</AccordionGroup>
</section>

<section class="actionButtonsContainer" aria-label="action buttons">
	<Button href="/" variant="outlined" color="secondary">cancel deployment</Button>
	<Button href="/initialize/review" variant="raised" color="secondary">review deployment</Button>
</section>
