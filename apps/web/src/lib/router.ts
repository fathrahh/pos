import type { FlowState, NavSection } from './types';

export interface RouteState {
  step: FlowState;
  nav: NavSection;
  redirectUrl?: string;
}

const VALID_NAVS: NavSection[] = ['overview', 'products', 'packages', 'vouchers', 'outlets'];

export function parseRoute(pathname: string, search: string = ''): RouteState {
  const cleanPath = pathname.replace(/^\/+|\/+$/g, '');
  const params = new URLSearchParams(search);
  const redirect = params.get('redirect') || undefined;

  // Root "/" defaults to dashboard overview
  if (!cleanPath) {
    return { step: 'DASHBOARD', nav: 'overview' };
  }

  // Explicit "/login"
  if (cleanPath === 'login') {
    return { step: 'LOGIN', nav: 'overview', redirectUrl: redirect };
  }

  const segments = cleanPath.split('/');
  if (segments[0] === 'dashboard') {
    const sub = segments[1] as NavSection;
    if (sub && VALID_NAVS.includes(sub)) {
      return { step: 'DASHBOARD', nav: sub };
    }
    return { step: 'DASHBOARD', nav: 'overview' };
  }

  // Fallback direct nav slug: e.g. /products
  if (VALID_NAVS.includes(segments[0] as NavSection)) {
    return { step: 'DASHBOARD', nav: segments[0] as NavSection };
  }

  // Unknown route falls back to dashboard overview
  return { step: 'DASHBOARD', nav: 'overview' };
}

export function buildPath(step: FlowState, nav: NavSection, redirectUrl?: string): string {
  if (step === 'LOGIN') {
    if (redirectUrl) {
      return `/login?redirect=${encodeURIComponent(redirectUrl)}`;
    }
    return '/login';
  }
  return `/dashboard/${nav}`;
}
