import type { FlowState, NavSection } from './types';

export interface RouteState {
  step: FlowState;
  nav: NavSection;
}

const VALID_NAVS: NavSection[] = ['overview', 'products', 'packages', 'vouchers', 'outlets'];

export function parseRoute(pathname: string): RouteState {
  const cleanPath = pathname.replace(/^\/+|\/+$/g, '');

  if (!cleanPath || cleanPath === 'login') {
    return { step: 'LOGIN', nav: 'overview' };
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

  return { step: 'LOGIN', nav: 'overview' };
}

export function buildPath(step: FlowState, nav: NavSection): string {
  if (step === 'LOGIN') {
    return '/login';
  }
  return `/dashboard/${nav}`;
}
