<?php
/**
 * Plugin Name: Container Badge
 * Description: Shows which container served the page. Baked into the image in Lab 04.
 * Version:     1.0.0
 */

add_action( 'wp_footer', function () {
	$host    = gethostname();
	$message = getenv( 'BADGE_MESSAGE' ) ?: 'Served by container';

	printf(
		'<div style="position:fixed;bottom:16px;right:16px;z-index:99999;display:flex;gap:8px;align-items:center;'
		. 'background:#0b1f33;color:#fff;font:600 14px/1.3 system-ui,sans-serif;padding:10px 14px;'
		. 'border-radius:12px;box-shadow:0 8px 24px rgba(0,0,0,.35)">'
		. '<span style="font-size:20px">🐳</span>%s '
		. '<code style="background:#13314f;color:#7dd3fc;padding:2px 8px;border-radius:6px;font-size:13px">%s</code></div>',
		esc_html( $message ),
		esc_html( $host )
	);
} );
